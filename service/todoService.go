package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"todo-app-go-fiber/dao"
	"todo-app-go-fiber/dto"
	"todo-app-go-fiber/entity"
)

//go:generate mockgen -destination=../mocks/service/mockTodoService.go -package=service todo-app-go-fiber/service TodoService
type DefaultTodoService struct {
	Repo dao.TodoRepository
}

type TodoService interface {
	TodoInsert(todo entity.Todo) (*dto.TodoDTO, error)
	TodoGetAll() ([]entity.Todo, error)
	TodoDelete(id primitive.ObjectID) (bool, error)
}

func (t DefaultTodoService) TodoInsert(todo entity.Todo) (*dto.TodoDTO, error) {
	var res dto.TodoDTO

	isTitleLengthLessThanThree := len(todo.Title) <= 3

	if isTitleLengthLessThanThree {
		res.Status = false
		return &res, nil
	}

	result, err := t.Repo.Insert(todo)

	isErrNotNil := err != nil
	isResultEqualsFalse := result == false

	if isErrNotNil || isResultEqualsFalse {
		res.Status = false
		return &res, err
	}

	todo.RequestId = dao.RandomRequestId
	res = dto.TodoDTO{Status: result, TimeStamp: time.Now(), Todo: todo}
	return &res, nil
}

func (t DefaultTodoService) TodoGetAll() ([]entity.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t DefaultTodoService) TodoDelete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)

	if err != nil || result == false {
		return false, err
	}

	return true, nil
}

func NewTodoService(Repo dao.TodoRepository) DefaultTodoService {
	return DefaultTodoService{Repo: Repo}
}
