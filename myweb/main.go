/**
 * @date: 2022/7/18
 * @desc:
 */

package main

import (
	"fmt"
	"net/http"
	"tea"
)

func imMiddleWare() tea.HandlerFunc {
	return func(context *tea.Context) {
		fmt.Println("我是中间件")
	}
}

func main() {
	r := tea.New()
	v1 := r.Group("/api")
	v1.Use(imMiddleWare())

	v1.GET("/group", func(c *tea.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *tea.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *tea.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *tea.Context) {
		c.JSON(http.StatusOK, tea.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}
