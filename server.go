package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type todo struct {
	id         int
	todo       string
	isComplete bool
}

func main() {
	allTodo := []*todo{}
	setTodo := map[string]bool{}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})

	r.POST("/add-todo", func(c *gin.Context) {
		title := c.PostForm("todo")

		if !setTodo[title] {
			setTodo[title] = true
			newId := time.Now().Nanosecond()
			tmp := &todo{newId, title, false}
			allTodo = append(allTodo, tmp)

			c.HTML(http.StatusOK, "newTodo.tmpl", gin.H{
				"todo":    tmp.todo,
				"id":      tmp.id,
				"checked": tmp.isComplete,
			})
		}
	})

	r.GET("/toggle-comp-state/:id", func(c *gin.Context) {
		idPar := c.Param("id")
		id, err := strconv.Atoi(idPar)

		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(allTodo); i++ {
			v := *allTodo[i]
			if v.id == id {
				v.isComplete = !v.isComplete

				allTodo[i] = &v

				fmt.Println(allTodo[i])

				c.HTML(http.StatusOK, "newTodo.tmpl", gin.H{
					"todo":    v.todo,
					"id":      v.id,
					"checked": v.isComplete,
				})
				return
			}
		}

	})

	r.Run(":3000")
}
