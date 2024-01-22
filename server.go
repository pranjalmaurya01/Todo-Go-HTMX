package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	id         int
	todo       string
	isComplete bool
}

func main() {
	allTodo := []todo{}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/add-todo", func(ctx *gin.Context) {
		tmp := todo{123, "asd", false}
		allTodo = append(allTodo, tmp)

		fmt.Println(allTodo)
	})

	r.Run(":3000")
}
