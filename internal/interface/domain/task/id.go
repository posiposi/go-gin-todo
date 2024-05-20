package task

import (
	"go-gin-todo/internal/interface/domain/common"
)

type ID struct {
	*common.UUID
}

func NewTodoId() (*ID, error) {
	id, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := ID{id}
	return &i, nil
}
