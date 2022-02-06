package main

import (
	"io"
	"net/http"
	"os"
	authservice "webREST/authService"
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
	server.Use(gin.Recovery(), loggingservice.Logger(), authservice.BasicAuth())

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	server.GET("/files", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, FileController.FindAll())
	})

	server.POST("/files", func(ctx *gin.Context) {
		err := FileController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Input payload is valid."})
		}
	})

	server.Run(":8080")
}
