package services

import (
	"API_Youtube/mocks/repository"
	"API_Youtube/models"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{primitive.NewObjectID(), "Name 1", "Content 1"},
	{primitive.NewObjectID(), "Name 1", "Content 1"},
}

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo = repository.NewMockTodoRepository(ctrl)
	service = NewTodoService(mockRepo)
	return func() {
		service = nil
		defer ctrl.Finish()
	}
}
