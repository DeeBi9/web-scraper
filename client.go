package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// With the status code equals to 200 it will parse the
// HTML content
func parseContent() {

}

func noContent() {
	//
}

func handleResponse(statusCode string) {
	code_list := strings.Split(statusCode, " ")
	code := code_list[0]

	switch {
	case code == "200": // The request is successful
		parseContent()
	case code == "204": // Server successfully processed the request but has no content
		noContent()
	case code == "301":
		//
	case code == "302":
		//
	case code == "304":
		//
	case code == "400": // Cannot process due to client error
		// Client Error
	case code == "401":
		// Unauthorized
	case code == "403":
		// Forbidden
	case code == "404":
		// Not Found
	}

}

func main() {
	// Specify the URL correctly, including the protocol and host.
	url := "http://localhost:8080"

	// Send the HTTP GET request and get the response.
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Print the response status and headers.
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Headers:", resp.Header)

	// Passing the status code to handleResponse
	handleResponse(resp.Status)

	bytes, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
