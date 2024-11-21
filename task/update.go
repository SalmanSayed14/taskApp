package task

import (
	"time"
)

func UpdateTask(id int, name, description string, deadline time.Time) {
	for i, t := range Tasks {
		if t.ID == id {
			Tasks[i].Name = name
			Tasks[i].Description = description
			Tasks[i].Deadline = deadline
			break
		}
	}
}
