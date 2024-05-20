package task

import (
	"errors"
	"time"
)

type DueDate struct {
	value time.Time
}

func NewDueDate(value time.Time) (*DueDate, error) {
	if value.IsZero() {
		return nil, errors.New("期限は必須です。")
	}
	if value.Before(time.Now()) {
		return nil, errors.New("期限は現在時刻より後の日時を指定してください。")
	}
	return &DueDate{value: value}, nil
}
