package main

import (
	"fmt"
	"github.com/Huiyicc/mhyapi/genshin"
)

func main() {
	test_Enka()
}

func test_Enka() {
	enka, err := genshin.NewEnkaCore("./cache")
	if err != nil {
		panic(err)
	}
	var rawMap map[string]any
	if rawMap, err = enka.GetUserDataRaw("112075042"); err != nil {
		panic(err)
	}
	fmt.Println(rawMap)
}
