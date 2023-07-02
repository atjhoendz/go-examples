package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(ctx *gin.Context) {
		user := ctx.Params.ByName("name")
		value, ok := db[user]

		if ok {
			ctx.JSON(http.StatusOK, gin.H{"user": user, "grade": value})
		} else {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		}
	})

	r.POST("/webhook", func(ctx *gin.Context) {
		var payload struct {
			Name  string `json:"name" binding:"required"`
			Grade string `json:"grade" binding:"required"`
		}

		if ctx.Bind(&payload) == nil {
			db[payload.Name] = payload.Grade
			ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in localhost:8080
	r.Run(":8080")
}
