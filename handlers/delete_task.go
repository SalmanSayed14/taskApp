package handlers

import (
	"net/http"
	"strconv"
	"taskApp/task"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	task.DeleteTask(id)

	// Redirect back to the index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
