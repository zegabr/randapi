package repository

import "randapi.com/entity"


type Repository interface {
    GetRow() (*entity.Row, error)
}
