package task

import "errors"

type Title struct {
	value string
}

func NewTitle(value string) (*Title, error) {
	if value == "" {
		return nil, errors.New("タイトルは必須です。")
	}
	if len(value) > 100 {
		return nil, errors.New("タイトルは100文字以内で入力してください。")
	}
	return &Title{value: value}, nil
}
