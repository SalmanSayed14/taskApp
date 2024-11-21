package handlers

import (
	"log"
	"net/http"
	"sort"
	"taskApp/task"
	"text/template"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	var tasksWithRemaining []struct {
		task.Task
		Remaining string
	}

	// Calculate remaining time for each task
	for _, t := range task.Tasks {
		remaining := task.CalculateRemainingTime(t.Deadline)
		tasksWithRemaining = append(tasksWithRemaining, struct {
			task.Task
			Remaining string
		}{
			Task:      t,
			Remaining: remaining,
		})
	}

	// Sort tasks by remaining time
	sort.Slice(tasksWithRemaining, func(i, j int) bool {
		durationI := task.CalculateRemainingDuration(tasksWithRemaining[i].Deadline)
		durationJ := task.CalculateRemainingDuration(tasksWithRemaining[j].Deadline)
		return durationI < durationJ
	})

	// Render the index template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing index template: %v", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Pass data to the template
	data := struct {
		Tasks []struct {
			task.Task
			Remaining string
		}
	}{Tasks: tasksWithRemaining}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error rendering index template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
