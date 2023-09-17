package dao

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
	"todo-app-go-fiber/entity"
)

var RandomRequestId primitive.ObjectID

//go:generate mockgen -destination=../mocks/dao/mockTodoRepository.go -package=dao todo-app-go-fiber/dao TodoRepository
type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo entity.Todo) (bool, error)
	GetAll() ([]entity.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo entity.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	RandomRequestId = primitive.NewObjectID()
	todo.RequestId = RandomRequestId

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	isIdNil := result.InsertedID == nil
	isHasError := err != nil

	if isIdNil || isHasError {
		errors.New("failed")
		return false, err
	}

	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]entity.Todo, error) {
	var todo entity.Todo
	var todos []entity.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"requestid": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
