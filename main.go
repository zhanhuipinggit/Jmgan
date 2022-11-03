package main

import (
	"fmt"
	"gan"
	"net/http"
)

func main() {

	engine := gan.New()
	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)

	})
	engine.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "key[%q]= %q\n", k, v)
		}
	})
	engine.Run(":9999")

}
