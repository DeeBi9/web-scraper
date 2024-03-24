package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://youtube.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
