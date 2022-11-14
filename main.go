package main

import (
	"gan"
	"net/http"
)

func main() {

	r := gan.New()
	r.GET("/", func(c *gan.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gan.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gan.Context) {
		// expect /hello/geektut

		c.String(http.StatusOK, "%s%s", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gan.Context) {
		c.Json(http.StatusOK, gan.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")

}
