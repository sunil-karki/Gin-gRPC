package main

import (
	"webfront/controller"
	"webfront/service"

	"github.com/gin-gonic/gin"
)

var (
	fileService    service.FileService       = service.New()
	FileController controller.FileController = controller.New(fileService)
)

func main() {
	server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "OK",
	// 	})
	// })

	server.GET("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, FileController.FindAll())
	})

	server.POST("/posts", func(ctx *gin.Context) {
		ctx.JSON(200, FileController.Save(ctx))
	})

	server.Run(":8080")
}
