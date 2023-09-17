package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"todo-app-go-fiber/entity"
	"todo-app-go-fiber/mocks/dao"
)

var mockRepo *dao.MockTodoRepository
var service TodoService

var FakeData = []entity.Todo{
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
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = dao.NewMockTodoRepository(ct)
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

func TestDefaultTodoService_TodoInsert(t *testing.T) {
	td := setup(t)
	defer td()

	todo := entity.Todo{
		RequestId: primitive.NewObjectID(),
		Title:     "Test Insert Title",
		Content:   "Test Insert Content",
	}

	mockRepo.EXPECT().Insert(todo).Return(true, nil)

	result, err := service.TodoInsert(todo)

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}

func TestDefaultTodoService_TodoDelete(t *testing.T) {
	td := setup(t)
	defer td()

	fakeId := FakeData[0].RequestId
	mockRepo.EXPECT().Delete(fakeId).Return(true, nil)

	result, err := service.TodoDelete(fakeId)

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
