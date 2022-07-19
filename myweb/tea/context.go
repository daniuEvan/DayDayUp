/**
 * @date: 2022/7/18
 * @desc: 封装context
 */

package tea

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request

	// request info
	Path   string
	Method string
	Params map[string]string

	// response info
	StatusCode int
}

// newContext is constructor of Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm get form element
func (context *Context) PostForm(key string) string {
	return context.Req.FormValue(key)
}

// Query get url args
func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}

// Status set response status code
func (context *Context) Status(code int) {
	context.StatusCode = code
	context.Writer.WriteHeader(code)
}

// SetHeader set header
func (context *Context) SetHeader(key string, value string) {
	context.Writer.Header().Set(key, value)
}

// Param 获取请求参数
func (context *Context) Param(key string) string {
	value, _ := context.Params[key]
	return value
}

// String response hold String
func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	context.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// JSON response hold json
func (context *Context) JSON(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.Writer, err.Error(), http.StatusBadRequest)
	}
}

// Data response hold byte data
func (context *Context) Data(code int, data []byte) {
	context.Status(code)
	context.Writer.Write(data)
}

// HTML response hold byte html
func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	context.Writer.Write([]byte(html))
}
