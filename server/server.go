// Mock server for testing the package
package main

import "net/http"

/*type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
type ResponseWriter interface {
	Header() Header
	Write([]byte) (int, error)
	WriteHeader(int)
}*/

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("./public/")))
}
