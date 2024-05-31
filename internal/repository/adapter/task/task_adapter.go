package task

import (
	"database/sql"
	"go-gin-todo/internal/domain/task"
	"go-gin-todo/internal/domain/user"
)

type TaskAdapter struct {
	db *sql.DB
}

func NewTaskAdapter(db *sql.DB) (*TaskAdapter, error) {
	return &TaskAdapter{db: db}, nil
}

func (ta *TaskAdapter) GetTasksByUserId(userId *user.UserId) ([]task.Task, error) {
	rows, err := ta.db.Query("SELECT * FROM tasks WHERE user_id = ?", userId.UUID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.Id, &t.Title, &t.Status, &t.DueDate, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
