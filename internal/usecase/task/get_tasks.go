package task

import (
	"go-gin-todo/internal/domain/task"
	port "go-gin-todo/internal/repository/port/task"
)

type GetTasks struct {
	taskPort port.TaskPort
}

func NewGetTasks(taskPort port.TaskPort) (*GetTasks, error) {
	return &GetTasks{taskPort: taskPort}, nil
}

// タスクが1件も存在しないかを判定する
func IsTasksExists(tasks []task.Task)
