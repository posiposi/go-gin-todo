package user

import (
	"go-gin-todo/internal/domain/common"
)

type UserId struct {
	*common.UUID
}

func NewUserId() (*UserId, error) {
	id, err := common.GenerateUUID()
	if err != nil {
		return nil, err
	}
	i := UserId{id}
	return &i, nil
}
