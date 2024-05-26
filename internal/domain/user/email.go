package user

import "errors"

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	// TODO 正規表現に変更する
	if value == "" {
		return nil, errors.New("氏名は必須です。")
	}
	if len(value) > MaxNameLength {
		return nil, errors.New("氏名は50文字以内で入力してください。")
	}
	return &Email{value: value}, nil
}
