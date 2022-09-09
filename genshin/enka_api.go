package genshin

import (
	"errors"
	"github.com/Huiyicc/mhyapi/tools"
	json "github.com/json-iterator/go"
	"io"
	"net/http"
	"os"
	"time"
)

var ErrorCacheIsNoteSet = errors.New("缓存不存在")

// NewEnkaCore 创建一个新的实例包
// cache为必填参数,因为EnkaNetwork有ttl,而且在境外
// 所以做cache是最好的
func NewEnkaCore(cachePath string) (*EnkaCore, error) {
	err := os.MkdirAll(cachePath, 655)
	if err != nil {
		return nil, err
	}
	lstr := cachePath[len(cachePath)-1:]
	if lstr != "/" && lstr != "\\" {
		cachePath += "/"
	}
	return &EnkaCore{
		cachePath: cachePath,
	}, nil
}

type EnkaCore struct {
	//缓存目录
	cachePath string
}

// GetUserDataRaw 使用UID获取玩家数据,返回原始键值对
func (t *EnkaCore) GetUserDataRaw(uid string) (map[string]any, error) {
	url := "https://enka.network/u/" + uid + "/__data.json"
	cachePath := t.cachePath + "uid_data"
	os.MkdirAll(cachePath, 0655)
	data, err := t.cacheHttpGet(url, cachePath)
	if err != nil {
		return nil, err
	}
	//解析数据
	var f map[string]any
	if err = json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	//获取ttl
	lstTime := f["ltime"]
	if lstTime == nil {
		//ttl不存在,是新的,设置ttl
		lstTime = time.Now().Unix()
		f["ltime"] = lstTime
		data, err = json.Marshal(f)
		//更新缓存
		if err = t.setCacheData(t.getUrlCachePath(url, cachePath), data); err != nil {
			return nil, err
		}
	}
	ltime, _ := lstTime.(float64)
	ttl, _ := f["ttl"].(float64)
	if float64(time.Now().Unix()) > ltime+ttl {
		//当ttl过期的时候刷新,删除缓存,重新请求
		if err = os.Remove(t.getUrlCachePath(url, cachePath)); err != nil {
			return nil, err
		}
		return t.GetUserDataRaw(uid)
	}
	return f, nil
}

func (t *EnkaCore) cacheHttpGet(url, cachePath string) ([]byte, error) {
	cp := t.getUrlCachePath(url, cachePath)
	if data, err := t.getCacheData(cp); err == nil && err != ErrorCacheIsNoteSet {
		return data, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request Status Code: " + resp.Status)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(cp, os.O_RDWR|os.O_CREATE, 0655)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t.setCacheData(cp, data)
	return data, err
}
func (t *EnkaCore) getCacheData(cachePath string) ([]byte, error) {
	if isSet, _ := tools.PathExists(cachePath); isSet {
		f, err := os.OpenFile(cachePath, os.O_RDWR, 0655)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return io.ReadAll(f)
	}
	return nil, ErrorCacheIsNoteSet
}
func (t *EnkaCore) setCacheData(cacheFileName string, data []byte) error {
	os.Remove(cacheFileName)
	f, err := os.OpenFile(cacheFileName, os.O_CREATE|os.O_RDWR, 0655)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}
func (t *EnkaCore) removeCache(cacheFileName string) error {
	return os.Remove(cacheFileName)
}
func (t *EnkaCore) getUrlCachePath(url, cachePath string) string {
	umd5 := tools.GetMd5(url)
	lstr := cachePath[len(cachePath)-1:]
	if lstr != "/" && lstr != "\\" {
		cachePath += "/"
	}
	return cachePath + umd5
}
