/**
 * @date: 2022/7/18
 * @desc: 路由功能基本实现
 */

package tea

import (
	"log"
	"net/http"
	"strings"
)

// router define router struct
// roots 使用 roots 来存储每种请求方式的Trie 树根节点。使用 handlers 存储每种请求方式的 HandlerFunc
type router struct {
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

// newRouter is constructor of router
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node), // key 是请求方法
	}
}

// parsePattern  only one * is allowed
func parsePattern(pattern string) (resParts []string) {
	parts := strings.Split(pattern, "/")
	for _, part := range parts {
		if part != "" {
			resParts = append(resParts, part)
			if strings.HasPrefix(part, "*") {
				break
			}
		}
	}
	return resParts
}

// addRoute 创建路由和处理函数逻辑
func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Router %4s - %s", method, pattern)
	parts := parsePattern(pattern)
	key := method + "-" + pattern

	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

// getRoute 解析了:和*两种匹配符的参数，返回一个 map
// 例如/p/go/doc匹配到/p/:lang/doc，解析结果为：{lang: "go"}， /static/css/geektutu.css匹
// 配到/static/*filepath，解析结果为{filepath: "css/geektutu.css"}
func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method] // 获取请求方法对于的前缀树

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0) // 获取节点node

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

// handle 根据上下文获取路由函数
func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
