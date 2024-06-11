package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
) 

type Test struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(ctx *gin.Context) {
		// ctx.String(http.StatusOK, "Hello World")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",	
			"message": "Hello World",
		})
	} )

	v1 := r.Group("/v1")
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",	
			"message": param,
		})
	})

	v1.POST("user", func(ctx *gin.Context) {
		var data Test

		ctx.BindJSON(&data)

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",	
			"message": data,
		})
	})

	v1.GET("/user/", func(ctx *gin.Context) {
		query := ctx.Query("name")

		ctx.JSON(http.StatusOK, gin.H{
			"status":  "success",	
			"message": query,
		})
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}