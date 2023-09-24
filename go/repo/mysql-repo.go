package repository

import "randapi.com/entity"

type repo struct{}

func NewMysqlRepository() Repository {
	return &repo{}
}

func (*repo) GetRow() (*entity.Row, error) {
	return &entity.Row{
		Language: "go",
		Phrase:   "hello from go", // TODO: connect to mysql and get prepopulated row for golang
	}, nil
}
