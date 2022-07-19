/**
 * @date: 2022/7/18
 * @desc:
 */

package tea

import (
	"log"
	"net/http"
)

// router define router struct
type router struct {
	handlers map[string]HandlerFunc
}

// newRouter is constructor of router
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

// addRoute 创建路由和处理函数逻辑
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle 根据上下文获取路由函数
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 not found: %s", c.Path)
	}
}
