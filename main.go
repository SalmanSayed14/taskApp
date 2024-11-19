package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Deadline    time.Time `json:"deadline"`
}

var tasks []Task
var taskID int
var mu sync.Mutex

// Calculate remaining time for each task
func calculateRemainingTime(deadline time.Time) string {
	now := time.Now()
	if deadline.Before(now) {
		return "Expired"
	}

	duration := deadline.Sub(now)
	days := int(duration.Hours()) / 24
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	// Format remaining time as "X days Y hours Z minutes"
	return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Create a new slice to hold tasks along with the remaining time and ShowDescription flag
	var tasksWithRemaining []struct {
		Task
		Remaining string
	}

	// Calculate remaining time for each task
	for _, task := range tasks {
		remaining := calculateRemainingTime(task.Deadline)
		tasksWithRemaining = append(tasksWithRemaining, struct {
			Task
			Remaining string
		}{
			Task:      task,
			Remaining: remaining,
		})
	}

	// Sort tasks based on remaining time (ascending order)
	sort.Slice(tasksWithRemaining, func(i, j int) bool {
		durationI := calculateRemainingDuration(tasksWithRemaining[i].Deadline)
		durationJ := calculateRemainingDuration(tasksWithRemaining[j].Deadline)
		return durationI < durationJ
	})

	// Load the template
	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		log.Printf("Error parsing index template: %v", err)
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Pass tasks with remaining time to the template
	data := struct {
		Tasks []struct {
			Task
			Remaining string
		}
	}{
		Tasks: tasksWithRemaining,
	}

	// Render the template
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error rendering index template: %v", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func calculateRemainingDuration(deadline time.Time) time.Duration {
	now := time.Now()
	if deadline.Before(now) {
		return -1 // Return a negative value if the task is expired
	}
	return deadline.Sub(now) // Calculate the remaining duration
}

// Serve the CSS file
func serveCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/styles.css")
}

// Add a new task
func addTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	name := r.PostFormValue("name")
	desc := r.PostFormValue("desc")
	deadlineStr := r.PostFormValue("deadline")

	if name == "" || deadlineStr == "" || desc == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Parse the deadline
	deadline, err := time.Parse("2006-01-02T15:04", deadlineStr)
	if err != nil {
		http.Error(w, "Invalid deadline format", http.StatusBadRequest)
		return
	}

	// Increment taskID and assign it to the new task
	taskID++
	newTask := Task{
		ID:          taskID,
		Name:        name,
		Description: desc,
		Deadline:    deadline,
	}

	tasks = append(tasks, newTask)

	// Redirect back to the index page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Delete a task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Extract the task ID from the URL path
	idStr := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Check for _method field to handle the "DELETE" method
	if r.Method == http.MethodPost && r.PostFormValue("_method") == "DELETE" {
		// Remove the task by ID
		for i, task := range tasks {
			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...)
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}
		}
		http.NotFound(w, r)
		return
	}

	// If no _method field or if it's not a POST, treat it as a 404
	http.NotFound(w, r)
}

func showTaskDetails(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Extract the task ID from the URL path
	idStr := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Error parsing task ID: %v", err) // Log the error for debugging
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Log that we're processing this request
	log.Printf("Fetching details for task ID: %d", taskID)

	// Find the task by ID
	var task *Task
	for _, t := range tasks {
		if t.ID == taskID {
			task = &t
			break
		}
	}

	if task == nil {
		log.Printf("Task with ID %d not found", taskID) // Log if task is not found
		http.NotFound(w, r)
		return
	}

	// Log the task data we're about to render
	log.Printf("Found task: %+v", task)

	// Parse the template file
	tmpl, err := template.ParseFiles("frontend/task_details.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err) // Log template parsing error
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Render the task details page by passing the task data to the template
	err = tmpl.Execute(w, task)
	if err != nil {
		log.Printf("Error rendering template: %v", err) // Log rendering error
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func showUpdateForm(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Extract the task ID from the URL path
	idStr := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Find the task by ID
	for _, task := range tasks {
		if task.ID == taskID {
			// Render the update form with the current task data
			tmpl, err := template.ParseFiles("frontend/update.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Pass the task data to the form
			err = tmpl.Execute(w, task)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	http.NotFound(w, r)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Extract the task ID from the URL path
	idStr := mux.Vars(r)["id"]
	taskID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	// Find the task by ID
	for i, task := range tasks {
		if task.ID == taskID {
			// Get the updated task data from the form
			name := r.PostFormValue("name")
			desc := r.PostFormValue("desc")
			deadlineStr := r.PostFormValue("deadline")

			if name == "" || desc == "" || deadlineStr == "" {
				http.Error(w, "Invalid input", http.StatusBadRequest)
				return
			}

			// Parse the deadline
			deadline, err := time.Parse("2006-01-02T15:04", deadlineStr)
			if err != nil {
				http.Error(w, "Invalid deadline format", http.StatusBadRequest)
				return
			}

			// Update the task
			tasks[i].Name = name
			tasks[i].Description = desc
			tasks[i].Deadline = deadline

			// Redirect back to the index page
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	// Serve static files (HTML and CSS)
	myRouter.HandleFunc("/", serveIndex)
	myRouter.HandleFunc("/styles.css", serveCSS)
	myRouter.PathPrefix("/sources/").Handler(http.StripPrefix("/sources", http.FileServer(http.Dir("./sources"))))

	// Define routes for adding tasks and deleting tasks
	myRouter.HandleFunc("/tasks", addTask).Methods("POST")         // Add a new task
	myRouter.HandleFunc("/tasks/{id}", deleteTask).Methods("POST") // Delete a task
	// Add route to show task details
	myRouter.HandleFunc("/tasks/{id}/details", showTaskDetails).Methods("GET")
	myRouter.HandleFunc("/tasks/{id}/update", showUpdateForm).Methods("GET") // Show update form
	myRouter.HandleFunc("/tasks/{id}/update", updateTask).Methods("POST")    // Process task update
	// Add route for toggling task description visibility

	// Start server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
