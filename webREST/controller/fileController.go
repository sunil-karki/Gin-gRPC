package controller

import (
	"webREST/model"
	"webREST/service"

	"github.com/gin-gonic/gin"
)

type FileController interface {
	FindAll() []model.File
	//To access data that comes with this context
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.FileService
}

func New(service service.FileService) FileController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []model.File {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var file model.File
	// ctx.BindJSON(&file)
	err := ctx.ShouldBindJSON(&file)
	if err != nil {
		return err
	}
	c.service.Save(file)
	return nil
}
