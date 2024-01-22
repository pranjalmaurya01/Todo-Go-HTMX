package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id         int
	Todo       string
	IsComplete bool
}

func main() {
	// Storing All Todo in memory
	allTodo := []*todo{}
	setTodo := map[string]bool{}

	// GIN CONFIG
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	// Loading All Todo in memory
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"data": allTodo,
		})
	})

	// Adding New Todo
	r.POST("/add-todo", func(c *gin.Context) {
		title := c.PostForm("todo")

		if !setTodo[title] {
			setTodo[title] = true
			newId := time.Now().Nanosecond()
			tmp := &todo{newId, title, false}
			allTodo = append(allTodo, tmp)

			c.HTML(http.StatusOK, "newTodo.tmpl", gin.H{
				"data": tmp,
			})
		}
	})

	// Toggling Complete State
	r.PATCH("/toggle-comp-state/:id", func(c *gin.Context) {
		idPar := c.Param("id")
		id, err := strconv.Atoi(idPar)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(allTodo); i++ {
			v := *allTodo[i]
			if v.Id == id {
				v.IsComplete = !v.IsComplete

				allTodo[i] = &v

				c.HTML(http.StatusOK, "newTodo.tmpl", gin.H{
					"data": v,
				})
				return
			}
		}

	})

	r.Run(":3000")
}
