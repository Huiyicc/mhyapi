package EnkaNetwork

import (
	_ "embed"
	"errors"
	json "github.com/json-iterator/go"
	"os"
)

const (
	NODE_URL_ENKA = "https://enka.network/"
	NODE_URL_CN1  = "https://enka.microgg.cn/"
	NODE_URL_CN2  = "https://enka.minigg.cn/"
)

var (
	ErrorCacheIsNoteSet = errors.New("缓存不存在")

	//go:embed characters.json
	characters []byte

	//go:embed dictionaries.json
	dictionariesRaw []byte

	dictionariesMap dictionaries
	charactersMap   charactersMapCore
)

func init() {
	if charactersMap == nil || len(charactersMap) == 0 {
		err := json.Unmarshal(characters, &charactersMap)
		if err != nil {
			panic(err)
		}
	}
	if !dictionariesMap.init {
		err := json.Unmarshal(dictionariesRaw, &dictionariesMap)
		if err != nil {
			panic(err)
		}
	}
}

// NewEnkaCore 创建一个新的实例包
// cache为必填参数,因为EnkaNetwork有ttl,而且在境外
// 所以做cache是最好的
func NewEnkaCore(cachePath, nodeUrl string) (*EnkaCore, error) {
	err := os.MkdirAll(cachePath, 655)
	if err != nil {
		return nil, err
	}
	lstr := cachePath[len(cachePath)-1:]
	if lstr != "/" && lstr != "\\" {
		cachePath += "/"
	}
	lstr = nodeUrl[len(nodeUrl)-1:]
	if lstr != "/" {
		cachePath += "/"
	}
	return &EnkaCore{
		cachePath: cachePath,
		nodeUrl:   nodeUrl,
	}, nil
}

type EnkaCore struct {
	//缓存目录
	cachePath string
	//节点网址
	nodeUrl string
}
