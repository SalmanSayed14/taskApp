package task

func AddTask(newTask Task) {
	taskID++
	newTask.ID = taskID
	Tasks = append(Tasks, newTask)
}
