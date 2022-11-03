package main

import (
	"gan"
	"net/http"
)

func main() {

	r := gan.New()
	r.GET("/", func(c *gan.Context) {
		c.Html(http.StatusOK, "<h1>Hello Gan</h1>")
	})
	r.GET("/hello", func(c *gan.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gan.Context) {
		c.Json(http.StatusOK, map[string]interface{}{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})

	r.Run(":9999")

}
