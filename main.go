package main

import (
	"fmt"
	"log"
	"net/http"
	"taskApp/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Serve static files (CSS)
	myRouter.HandleFunc("/styles.css", handlers.ServeCSS)
	myRouter.PathPrefix("/sources/").Handler(http.StripPrefix("/sources", http.FileServer(http.Dir("./sources"))))

	// Define routes for tasks
	myRouter.HandleFunc("/", handlers.ServeIndex)                           // Serve tasks index
	myRouter.HandleFunc("/tasks", handlers.AddTask).Methods("POST")         // Add task
	myRouter.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("POST") // Delete task
	myRouter.HandleFunc("/tasks/{id}/details", handlers.ShowTaskDetails).Methods("GET")
	myRouter.HandleFunc("/tasks/{id}/update", handlers.ShowUpdateForm).Methods("GET") // Show update form
	myRouter.HandleFunc("/tasks/{id}/update", handlers.UpdateTask).Methods("POST")    // Process task update

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
