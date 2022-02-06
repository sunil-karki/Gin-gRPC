package service

import "webREST/model"

type FileService interface {
	Save(model.File) model.File
	FindAll() []model.File
}

type fileService struct {
	files []model.File
}

func New() FileService {
	return &fileService{}
}

func (service *fileService) Save(file model.File) model.File {
	service.files = append(service.files, file)
	return file
}

func (service *fileService) FindAll() []model.File {
	return service.files
}
