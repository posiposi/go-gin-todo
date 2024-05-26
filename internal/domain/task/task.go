package task

import "time"

type Task struct {
	Id        *Id
	Title     *Title
	Status    *Status
	DueDate   *DueDate
	CreatedAt time.Time
	UpdatedAt time.Time
}
