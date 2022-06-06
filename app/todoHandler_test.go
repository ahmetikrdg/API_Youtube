package app

import (
	services "API_Youtube/mocks/service"
	"API_Youtube/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
)

var ch TodoHandler

/*
func setup(t *testing.T) func() {
	ct := gomock.NewController(t)

	mockService = services.NewMockTodoService(ct)
	err := TodoHandler{mockService}
	if err.Service != nil {
		log.Fatalln(false)
	}
	return func() {
		defer ct.Finish()
	}

---
//arrange
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}*/

func TestTodoHandler_GetAllTodo(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockService := services.NewMockTodoService(ctrl)
	td := TodoHandler{mockService}

	router := fiber.New()
	router.Get("/api/todos", td.GetAllTodo)

	var FakeDataForHandler = []models.Todo{
		{primitive.NewObjectID(), "Title 1", "Content 1"},
		{primitive.NewObjectID(), "Title 2", "Content 2"},
		{primitive.NewObjectID(), "Title 3", "Content 3"},
	}
	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	resp, _ := router.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
