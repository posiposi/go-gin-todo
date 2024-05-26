package task

import (
	"go-gin-todo/internal/domain/common"
)

type Id struct {
	*common.UUID
}

func NewTodoId() (*Id, error) {
	id, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := Id{id}
	return &i, nil
}
