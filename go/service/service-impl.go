package service

import (
	"randapi.com/entity"
	repository "randapi.com/repo"
)

type service struct {}

func NewService() Service{
    return &service{}
}

var (
	repo repository.Repository = repository.NewMysqlRepository()
)

func (*service) GetRow() (*entity.Row, error) {
	return repo.GetRow()
}

