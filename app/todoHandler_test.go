package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http/httptest"
	"testing"
	"todo-app-go-fiber/entity"
	"todo-app-go-fiber/mocks/service"
)

var todo TodoHandler
var mockService *service.MockTodoService

var HandlerFakeData = []entity.Todo{
	{
		RequestId: primitive.NewObjectID(),
		Title:     "Test Title 1",
		Content:   "Test Content 1",
	},
	{
		RequestId: primitive.NewObjectID(),
		Title:     "Test Title 2",
		Content:   "Test Content 2",
	},
	{
		RequestId: primitive.NewObjectID(),
		Title:     "Test Title 3",
		Content:   "Test Content 3",
	},
}

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mockService = service.NewMockTodoService(ctrl)

	todo = TodoHandler{mockService}

	return func() {
		defer ctrl.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {

	trd := setup(t)
	defer trd()

	router := fiber.New()
	router.Get("/api/todos", todo.GetAllTodo)

	mockService.EXPECT().TodoGetAll().Return(HandlerFakeData, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	response, _ := router.Test(req, 1)
	assert.Equal(t, 200, response.StatusCode)
}
