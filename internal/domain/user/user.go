package user

import "time"

type User struct {
	UserId    *UserId
	Name      *Name
	Password  *Password
	Email     *Email
	CreatedAt time.Time
	UpdatedAt time.Time
}
