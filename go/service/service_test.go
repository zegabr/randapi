package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"randapi.com/entity"
)

type MockRepo struct {
    mock.Mock
}

func (mock *MockRepo) GetRow() (*entity.Row, error) {
    args := mock.Called()
    result := args.Get(0)
    return result.(*entity.Row), args.Error(1)
}

func TestGetRow(t *testing.T){
    mockRepo := new(MockRepo)
    row := entity.Row{Language: "go", Phrase: "hello from go"}
    mockRepo.On("GetRow").Return(&row, nil)

    testService := NewService(mockRepo)
    result, _ := testService.GetRow()

    mockRepo.AssertExpectations(t)

    assert.Equal(t, result.Language, "go")
    assert.Equal(t, result.Phrase, "hello from go")
}
