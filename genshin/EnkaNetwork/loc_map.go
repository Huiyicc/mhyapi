package EnkaNetwork

import (
	"errors"
	"strconv"
)

const (
	HASHTEXT_LANGUAGE_PORTUGUESA    = "pt"    //葡萄牙语
	HASHTEXT_LANGUAGE_KOREA         = "ko"    //韩语
	HASHTEXT_LANGUAGE_JAMAICA       = "ja"    //牙买加
	HASHTEXT_LANGUAGE_INDONESIA     = "id"    //印度尼西亚
	HASHTEXT_LANGUAGE_ZHTW          = "zh-TW" //繁体中文
	HASHTEXT_LANGUAGE_RUSSIA        = "ru"    //俄罗斯
	HASHTEXT_LANGUAGE_VIRGINISLANDS = "vi"    //美属维尔京群岛
	HASHTEXT_LANGUAGE_THAILAND      = "th"    //泰语
	HASHTEXT_LANGUAGE_ZHCN          = "zh-CN" //简体中文
	HASHTEXT_LANGUAGE_DEUTSCHLAND   = "de"    //德语
	HASHTEXT_LANGUAGE_ENGLISH       = "en"    //英语
	HASHTEXT_LANGUAGE_FRANCE        = "fr"    //法国
	HASHTEXT_LANGUAGE_ESPANA        = "es"    //西班牙
)

type locTextMapCore map[string]map[string]string

// GetHashStr 用于将hash值转为文本
func GetHashStr(lang string, hash int64) (string, error) {
	str, isSet := locTextMap[lang][strconv.FormatInt(hash, 10)]
	if !isSet {
		return "", errors.New("hash不存在")
	}
	return str, nil
}
