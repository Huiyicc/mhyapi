package EnkaNetwork

import (
	"errors"
	"github.com/Huiyicc/mhyapi/tools"
	"io"
	"net/http"
	"os"
)

func (t *EnkaCore) httpGetIncidentalCache(url, cacheFilename string) ([]byte, error) {
	if data, err := t.getCacheData(cacheFilename); err == nil && err != ErrorCacheIsNoteSet {
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
	f, err := os.OpenFile(cacheFilename, os.O_RDWR|os.O_CREATE, 0655)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	t.setCacheData(cacheFilename, data)
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
