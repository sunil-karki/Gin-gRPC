package main

import (
	"io"
	"os"
	"webREST/controller"
	loggingservice "webREST/loggingService"
	"webREST/service"

	"github.com/gin-gonic/gin"
)

var (
	fileService    service.FileService       = service.New()
	FileController controller.FileController = controller.New(fileService)
)

func setupLogOutput() {
	f, _ := os.Create("application.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery(), loggingservice.Logger())

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	server.GET("/files", func(ctx *gin.Context) {
		ctx.JSON(200, FileController.FindAll())
	})

	server.POST("/files", func(ctx *gin.Context) {
		ctx.JSON(200, FileController.Save(ctx))
	})

	server.Run(":8080")
}
