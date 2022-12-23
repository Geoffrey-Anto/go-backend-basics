package main

import (
	"github.com/gin-gonic/gin"

	"strconv"
)

var todos = []map[string]string{}

func GetAllTodos(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}

func GetATodo(ctx *gin.Context) {
	var currTodos = []map[string]string{}
	name := ctx.Query("name")

	for i := 0; i < len(todos); i++ {
		if name == todos[i]["name"] {
			currTodos = append(currTodos, todos[i])
		}
	}

	if len(currTodos) == 0 {
		ctx.JSON(400, gin.H{
			"message": "No todos found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"todo": currTodos[0]["todo"],
	})
}

func AddTodo(ctx *gin.Context) {
	res := ctx.Query("todo")

	isPresent := false

	for i := 0; i < len(todos); i++ {
		if res == todos[i]["todo"] {
			isPresent = true
		}
	}

	if isPresent {
		ctx.JSON(400, gin.H{
			"message": "Todo Already Exists",
		})
		return
	}

	var length = len(todos)

	var str = strconv.Itoa(length)

	todo := map[string]string{
		"name": "Todo" + str,
		"todo": res,
	}

	todos = append(todos, todo)

	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}

func DeleteTodo(ctx *gin.Context) {
	name := ctx.Query("name")

	if name == "" {
		ctx.JSON(400, gin.H{
			"message": "Please provide a name",
		})
		return
	}

	// delete the todo
	for i := 0; i < len(todos); i++ {
		if name == todos[i]["name"] {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}

func DeleteAllTodos(ctx *gin.Context) {
	todos = []map[string]string{}

	ctx.JSON(200, gin.H{
		"todos": todos,
	})
}

func main() {
	r := gin.Default()
	r.GET("/todos", GetAllTodos)
	r.GET("/todo", GetATodo)
	r.POST("/add-todo", AddTodo)
	r.DELETE("/delete-todo", DeleteTodo)
	r.DELETE("/delete-todos", DeleteAllTodos)
	r.Run("127.0.0.1:4000")
}
