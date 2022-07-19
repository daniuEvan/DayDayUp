/**
 * @date: 2022/7/18
 * @desc:
 */

package tea

import (
	"net/http"
)

// HandlerFunc interface defines
type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// addRouter 路由添加
func (engine *Engine) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	engine.router.addRouter(method, pattern, handlerFunc)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("GET", pattern, handlerFunc)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handlerFunc HandlerFunc) {
	engine.addRouter("POST", pattern, handlerFunc)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	err = http.ListenAndServe(addr, engine)
	return err
}

// ServeHTTP defines
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	context := newContext(w, req)
	engine.router.handle(context)
}
