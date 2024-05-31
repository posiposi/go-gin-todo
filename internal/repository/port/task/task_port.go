package task

import (
	"go-gin-todo/internal/domain/task"
	"go-gin-todo/internal/domain/user"
)

type TaskPort interface {
	GetTasksByUserId(userId *user.UserId) ([]task.Task, error)
}
