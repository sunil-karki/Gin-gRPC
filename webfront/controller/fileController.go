package controller

import (
	"webfront/entity"
	"webfront/service"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	FindAll() []entity.File
	//To access data that comes with this context
	Save(ctx *gin.Context) entity.File
}

type controller struct {
	service service.FileService
}

func New(service service.FileService) FileController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.File {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.File {
	var file entity.File
	ctx.BindJSON(&file)
	c.service.Save(file)
	return file
}
