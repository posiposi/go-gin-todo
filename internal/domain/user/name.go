package user

import "errors"

const (
	MaxNameLength = 50
)

type Name struct {
	value string
}

func NewTitle(value string) (*Name, error) {
	if value == "" {
		return nil, errors.New("氏名は必須です。")
	}
	if len(value) > MaxNameLength {
		return nil, errors.New("氏名は50文字以内で入力してください。")
	}
	return &Name{value: value}, nil
}
