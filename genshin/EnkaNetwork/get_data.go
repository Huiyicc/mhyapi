package EnkaNetwork

import (
	json "github.com/json-iterator/go"
	"os"
	"time"
)

// GetUserDataRaw 使用UID获取玩家数据,返回原始键值对
func (t *EnkaCore) GetUserDataRaw(uid string, ttl ...bool) (CharactersCore, error) {
	url := t.nodeUrl + "u/" + uid + "/__data.json"
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
	if ltime == 0 {
		lint64, _ := lstTime.(int64)
		ltime = float64(lint64)
	}
	var ttlb float64
	ttlb, _ = f["ttl"].(float64)
	if len(ttl) > 0 {
		if !ttl[0] {
			ttlb = ltime
		}
	}
	if float64(time.Now().Unix()) > ltime+ttlb {
		//当ttl过期的时候刷新,删除缓存,重新请求
		if err = os.Remove(t.getUrlCachePath(url, cachePath)); err != nil {
			return nil, err
		}
		return t.GetUserDataRaw(uid)
	}
	return f, nil
}
