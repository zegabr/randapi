package service

import "randapi.com/entity"

type Service interface {
	GetRow() (*entity.Row, error)
}
