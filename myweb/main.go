/**
 * @date: 2022/7/18
 * @desc:
 */

package main

import (
	"net/http"
	"tea"
)

func HelloWord(ctx *tea.Context) {
	ctx.JSON(http.StatusOK, tea.H{"data": "hello word"})
}

func main() {
	router := tea.New()
	router.GET("/hello", HelloWord)
	router.Run(":8080")
}
