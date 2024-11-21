package handlers

import (
	"net/http"
	"strconv"
	"taskApp/task"
	"text/template"

	"github.com/gorilla/mux"
)

func ShowTaskDetails(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Find the task by ID
	t := task.GetTaskByID(id)
	if t == nil {
		http.NotFound(w, r)
		return
	}

	// Render the task details page
	tmpl, err := template.ParseFiles("templates/task_details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
