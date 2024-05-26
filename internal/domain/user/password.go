package user

import "errors"

const (
	MaxPasswordLength = 8
)

type Password struct {
	value string
}

func NewPassword(value string) (*Password, error) {
	// TODO 正規表現に変更する
	if value == "" {
		return nil, errors.New("パスワードは必須です。")
	}
	if len(value) <= MaxPasswordLength {
		return nil, errors.New("パスワードは8文字以下で入力してください。")
	}
	return &Password{value: value}, nil
}
