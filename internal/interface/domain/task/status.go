package task

import "errors"

const (
	Todo  = "todo"
	Doing = "doing"
	Done  = "done"
)

var validStatus = map[string]bool{
	Todo:  true,
	Doing: true,
	Done:  true,
}

type Status struct {
	value string
}

func NewStatus(value string) (*Status, error) {
	if value == "" {
		return nil, errors.New("ステータスは必須です。")
	}
	if _, ok := validStatus[value]; !ok {
		return nil, errors.New("不正なステータスです。")
	}
	return &Status{value: value}, nil
}
