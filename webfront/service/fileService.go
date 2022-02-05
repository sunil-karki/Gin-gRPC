package service

import "webfront/entity"

type FileService interface {
	Save(entity.File) entity.File
	FindAll() []entity.File
}

type fileService struct {
	files []entity.File
}

func New() FileService {
	return &fileService{}
}

func (service *fileService) Save(file entity.File) entity.File {
	service.files = append(service.files, file)
	return file
}

func (service *fileService) FindAll() []entity.File {
	return service.files
}
