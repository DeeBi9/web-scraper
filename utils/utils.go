package utils

import (
	"fmt"
	"net/http"
)

var url string

func SpecifyUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Incorrect url")
		panic(err)
	}
	fmt.Println(resp)
}

type utils struct {
	id    string
	class string
}

func UseId() {
	//
}

func UsePath() {
	//
}

func UsePath() {
	//
}

func UseCssSelector() {
	//
}

func TagName() {
	//
}

func UseLinkText() {
	//
}

func UseName() {
	//
}

func UseAttr() {
	//
}
