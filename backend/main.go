package main

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todoList = []Todo{}

func getTodoList(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

func postTodo(c *gin.Context) {
	var newTodo Todo
	c.BindJSON(&newTodo)

	newTodo.ID = strconv.Itoa(len(todoList) + 1)
	todoList = append(todoList, newTodo)

	c.JSON(http.StatusOK, newTodo)
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, t := range todoList {
		if t.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}
	r.Use(cors.New(config))

	r.GET("/todo", getTodoList)
	r.POST("/todo", postTodo)
	r.DELETE("/todo/:id", deleteTodo)

	r.Run(":8080")
}
