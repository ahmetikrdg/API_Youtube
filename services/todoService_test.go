package services

import (
	"API_Youtube/mocks/repository"
	"API_Youtube/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Title 1", "Content 1"},
	{primitive.NewObjectID(), "Title 2", "Content 2"},
	{primitive.NewObjectID(), "Title 3", "Content 3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)
	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)
	result, err := service.TodoGetAll()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
