package service

import (
	"randapi.com/entity"
	repository "randapi.com/repo"
)

type service struct {}

var (
    repositoryInstance repository.Repository
)

func NewService(repos repository.Repository) Service{
    repositoryInstance = repos
    return &service{}
}

func (*service) GetRow() (*entity.Row, error) {
	return repositoryInstance.GetRow()
}

