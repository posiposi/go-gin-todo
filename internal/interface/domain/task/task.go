package task

import "time"

type Task struct {
	ID        *ID
	Title     *Title
	Status    *Status
	DueDate   *DueDate
	CreatedAt time.Time
	UpdatedAt time.Time
}
