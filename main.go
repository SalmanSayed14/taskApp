package main

import (
	"fmt"
	"log"
	"net/http"
	"taskApp/handlers"

	"github.com/gorilla/mux"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/styles.css", handlers.ServeCSS)
	myRouter.PathPrefix("/sources/").Handler(http.StripPrefix("/sources", http.FileServer(http.Dir("./sources"))))

	myRouter.HandleFunc("/", handlers.ServeIndex)
	myRouter.HandleFunc("/tasks", handlers.AddTask).Methods("POST")
	myRouter.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("POST")
	myRouter.HandleFunc("/tasks/{id}/details", handlers.ShowTaskDetails).Methods("GET")
	myRouter.HandleFunc("/tasks/{id}/update", handlers.ShowUpdateForm).Methods("GET")
	myRouter.HandleFunc("/tasks/{id}/update", handlers.UpdateTask).Methods("POST")

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}
