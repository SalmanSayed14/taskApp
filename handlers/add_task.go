package handlers

import (
	"net/http"
	"taskApp/task"
	"time"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	desc := r.PostFormValue("desc")
	deadlineStr := r.PostFormValue("deadline")

	if name == "" || deadlineStr == "" || desc == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	deadline, err := time.Parse("2006-01-02T15:04", deadlineStr)
	if err != nil {
		http.Error(w, "Invalid deadline format", http.StatusBadRequest)
		return
	}

	newTask := task.Task{
		Name:        name,
		Description: desc,
		Deadline:    deadline,
	}

	task.AddTask(newTask)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
