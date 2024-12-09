package task

import (
	"time"
)

var Tasks []Task
var taskID int

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Deadline    time.Time `json:"deadline"`
}

func GetTaskByID(id int) *Task {
	for _, t := range Tasks {
		if t.ID == id {
			return &t
		}
	}
	return nil
}
