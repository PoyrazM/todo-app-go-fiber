package dto

import (
	"time"
	"todo-app-go-fiber/entity"
)

type TodoDTO struct {
	Status    bool        `json:"status,omitempty"`
	TimeStamp time.Time   `json:"time_stamp"`
	Todo      entity.Todo `json:"todos"`
}
