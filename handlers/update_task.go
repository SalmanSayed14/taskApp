package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"taskApp/task"
	"time"

	"github.com/gorilla/mux"
)

func ShowUpdateForm(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	t := task.GetTaskByID(id)
	if t == nil {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/update.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	name := r.PostFormValue("name")
	desc := r.PostFormValue("desc")
	deadlineStr := r.PostFormValue("deadline")

	deadline, err := time.Parse("2006-01-02T15:04", deadlineStr)
	if err != nil {
		http.Error(w, "Invalid deadline format", http.StatusBadRequest)
		return
	}

	task.UpdateTask(id, name, desc, deadline)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
