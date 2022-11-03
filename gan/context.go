package gan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Path    string
	Method  string
	Code    int
}

func CreateContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
		Path:    r.URL.Path,
		Method:  r.Method,
	}
}

func (c *Context) PostFrom(key string) string {
	return c.Request.FormValue(key)
}

func (c *Context) SetStatus(code int) {
	c.Code = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置头部
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 字符串返回
func (c *Context) String(code int, format string, value ...interface{}) {
	c.SetStatus(code)
	c.SetHeader("Content-Type", "text/plain")
	c.Writer.Write([]byte(fmt.Sprintf(format, value...)))
}

func (c *Context) Json(code int, object interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.SetStatus(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(object); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

func (c *Context) Html(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.SetStatus(code)
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte) {
	c.SetStatus(code)
	c.Writer.Write(data)
}
