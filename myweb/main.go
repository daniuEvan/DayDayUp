/**
 * @date: 2022/7/18
 * @desc:
 */

package main

import (
	"net/http"
	"tea"
)

func main() {
	r := tea.New()
	r.GET("/", func(c *tea.Context) {
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
