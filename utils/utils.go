package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type utils struct {
	id    string
	class string
}
type store struct {
	url  string
	resp *http.Response
}

// Object of store datatype
var st store

// This is function to get response for the given url by the user
func (st store) getResponse() *http.Response {
	response := st.resp

	code_list := strings.Split(response.Status, " ")
	code := code_list[0]

	switch {
	case code == "200": // The request is successful
		fmt.Println("TRUE")
	case code == "400": // Cannot process due to client error
		fmt.Println("Client Error")
	case code == "401":
		fmt.Println("Unauthorized")
	case code == "404":
		fmt.Println("Page Not Found")
	}
	return response
}

// This function is to set the URL entered by the user.
func (st store) Url(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Incorrect url")
		panic(err)
	}
	st.resp = resp
	st.url = url
}

func parseContent(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	tokenizer := html.NewTokenizer(strings.NewReader(string(body)))
	for {
		tokenType := tokenizer.Next() // Increment the next token type
		token := tokenizer.Token()    // Increment the next token value

		/* Token type error or argument error handler */
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				return
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return
		}
		fmt.Printf("Token: %v\n", html.UnescapeString(token.String()))
	}
}

func getHtml() {
	resp := st.resp
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(string(body))
}

func UseTag() {
	//
}

func UseId() {
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
