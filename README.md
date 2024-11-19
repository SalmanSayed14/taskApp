## TaskApp - Task Management Web Application

TaskApp is a simple web-based task management application that allows users to add, update, view, and delete tasks. It helps manage tasks with deadlines and displays the remaining time for each task.

# Features

- Task List: View all tasks sorted by the remaining time to the deadline.
- Add Task: Add new tasks with a name, description, and deadline.
- View Task Details: View detailed information about a specific task, including its name, description, and deadline.
- Update Task: Edit an existing task's name, description, and deadline.
- Delete Task: Remove tasks from the list.
  Responsive Design: The application is fully responsive, providing a smooth experience on both desktop and mobile devices.

# Technologies Used

- Go (Golang): The backend is built using Go, a fast and efficient programming language.
- Gorilla Mux: A powerful HTTP router and URL matcher for Go.
- HTML5, CSS3: The frontend is built using standard web technologies (HTML, CSS).
- Go Templates: Used for rendering dynamic content on the frontend.

- Project Structure
  /taskapp
  ├── main.go # The main Go program with routes and logic
  ├── frontend
  │ ├── index.html # The main page with the task list and add form
  │ ├── task_details.html # Template for task detail page
  │ ├── update.html # Template for updating a task
  │ └── styles.css # Styles for the app
  ├── sources # Directory for additional files (like images, etc.)
  └── README.md # Project documentation

# Installation

# Prerequisites

- Install Go (version 1.16 or higher).
- Install Gorilla Mux:

go get -u github.com/gorilla/mux

# Steps to Run Locally

Clone this repository:

git clone https://github.com/yourusername/taskapp.git
cd taskapp

Install dependencies:

go get -u github.com/gorilla/mux

Run the Go server:

go run main.go

Open a browser and navigate to http://localhost:8080 to start using the application.

# Endpoints

- GET /: Displays the task list and the form to add a new task.
- POST /tasks: Adds a new task with the provided name, description, and deadline.
- GET /tasks/{id}/details: Shows detailed information about the task.
- GET /tasks/{id}/update: Displays a form to update the task.
- POST /tasks/{id}/update: Updates the task with new information.
- POST /tasks/{id}: Deletes a task.

# Example Routes

- Add task: POST /tasks
- View task details: GET /tasks/{id}/details
- Update task: POST /tasks/{id}/update
- Delete task: POST /tasks/{id}

# Templates

- index.html - Displays the task list and the "Add New Task" form.
- task_details.html - Displays the details of a specific task.
- update.html - A form for updating an existing task's details.
- styles.css - The stylesheet for the app's look and feel.

# Contributions

- Feel free to fork this repository, submit issues, or create pull requests. Contributions are welcome!

# Future Improvements

- User Authentication: Implement user authentication to allow multiple users to manage their own tasks.
- Task Priority: Add functionality for setting priorities on tasks.
- Task Categories: Group tasks into categories to better organize them.
- Email Notifications: Add email notifications for upcoming deadlines.
