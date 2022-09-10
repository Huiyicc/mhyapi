package EnkaNetwork

import (
	json "github.com/json-iterator/go"
	"os"
	"time"
)

// GetUserDataRaw 使用UID获取玩家数据,返回原始键值对
func (t *EnkaCore) GetUserDataRaw(uid string, ttl ...bool) (*CharactersCore, error) {
	url := nodeUrl + "u/" + uid + "/__data.json"
	cachePath := t.cachePath + "uid_data"
	os.MkdirAll(cachePath, 0655)
	data, err := t.cacheHttpGet(url, cachePath)
	if err != nil {
		return nil, err
	}
	//解析数据
	var f CharactersCore
	if err = json.Unmarshal(data, &f); err != nil {
		return nil, err
	}
	//获取ttl
	lstTime := f.Ltime
	if lstTime == 0 {
		//ttl=0,是新的,设置ttl
		lstTime = time.Now().Unix()
		f.Ltime = lstTime
		data, err = json.Marshal(f)
		//更新缓存
		if err = t.setCacheData(t.getUrlCachePath(url, cachePath), data); err != nil {
			return nil, err
		}
	}
	ttlb := int64(f.Ttl)
	if len(ttl) > 0 {
		if !ttl[0] {
			ttlb = lstTime
		}
	}
	if time.Now().Unix() > lstTime+ttlb {
		//当ttl过期的时候刷新,删除缓存,重新请求
		if err = os.Remove(t.getUrlCachePath(url, cachePath)); err != nil {
			return nil, err
		}
		return t.GetUserDataRaw(uid)
	}
	return &f, nil
}
