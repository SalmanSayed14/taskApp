## TaskApp - A Simple Task Management Application

TaskApp is a simple task management web application built in Go. It allows users to add, update, delete, and view tasks. Each task has a name, description, and a deadline. The application displays the remaining time until each task's deadline and allows users to update task details or delete them.

# Features

- Add new tasks with a name, description, and deadline.
- View tasks with their remaining time until the deadline.
- Update task details, including name, description, and deadline.
- Delete tasks from the list.
- Clean, responsive design for managing tasks.

# Installation

Prerequisites

- Go (v1.18 or higher)
- A web browser for accessing the application.

# Steps to Run the Project

1- Clone the repository:
git clone https://github.com/yourusername/taskapp.git
cd taskapp

2- Install dependencies: The project uses gorilla/mux for routing. Install it by running:
go get github.com/gorilla/mux

3- Run the application: After installing dependencies, run the application using the following command:
go run main.go

4- Access the application: Open your browser and go to http://localhost:8080 to start managing tasks.

# Project Structure

TaskApp/

- ├── handlers/ # Handlers for HTTP routes
- │ ├── add_task.go # Add new task
- │ ├── delete_task.go # Delete task
- │ ├── update_tsk_details.go # Show update form
- │ ├── update.go # Update task details
- │ ├── show_task_details.go # Show task details
- │ ├── serve_css.go # Serve CSS file
- │ └── index.go # Render task list on homepage
- ├── task/ # Task-related logic
- │ ├── create.go # Add new task to the list
- │ ├── delete.go # Delete a task
- │ ├── update.go # Update task
- │ ├── calculate.go # Calculate remaining time for tasks
- │ └── task.go # Task struct and global task list
- ├── templates/ # HTML templates
- │ ├── index.html # Homepage - Task list
- │ ├── task_details.html # Task details view
- │ └── update.html # Task update form
- ├── frontend/ # Static assets
- │ └── styles.css # Stylesheet for the app
- ├── main.go # Main entry point for the application
- └── README.md # Project documentation (this file)

# How it Works

- Task List: The homepage shows a list of tasks. For each task, the remaining time until the deadline is displayed. Users can view details, update, or delete tasks.
- Adding Tasks: Users can add tasks via the form on the homepage by providing a name, description, and deadline.
- Updating Tasks: Tasks can be updated by clicking the "Update" button next to a task. Users are presented with a form to modify the task's details.
- Task Details: Clicking the "Details" button shows more information about the task, including the full description and deadline.
- Deleting Tasks: Tasks can be deleted by clicking the "Delete" button next to the task.

# Utilities

- utils.go: Contains helper functions such as error logging, date parsing, and formatting the remaining time.
- calculate.go: Provides functions to calculate the remaining time for each task.

# Running Tests

There are no built-in tests in the current version, but you can manually test the following features:

- Add tasks with various deadlines and check if they appear correctly on the homepage.
- Update task details and verify the changes.
  Delete tasks and ensure they no longer appear on the task list.

# Future Improvements

- Add user authentication to allow multiple users to manage their tasks independently.
- Store tasks in a database (currently tasks are stored in memory).
- Implement task priority levels and filtering options.
- Create an API to interact with the task list programmatically (RESTful API).
