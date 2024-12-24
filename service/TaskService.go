package service

import (
	"database/sql"
	"log"
	"time"

	"github.com/andreparelho/task-manager-golang/model"
)

type TaskService struct {
	Database    *sql.DB
	TaskChannel chan model.Task
}

func TaskConstructor(database *sql.DB, channel chan model.Task) *TaskService {
	return &TaskService{
		Database:    database,
		TaskChannel: channel,
	}
}

func (taskService *TaskService) AddTask(task model.Task) error {
	var query string = "INSERT INTO tasks (title, description, status, created_at) VALUES (?, ?, ?, ?)"
	_, errorCreate := taskService.Database.Exec(query, task.Title, task.Description, task.Status, task.CreatedAt)
	return errorCreate
}

func (taskService *TaskService) UpdateTask(task *model.Task) error {
	var query string = "UPDATE tasks SET status = ? WHERE id = ?"
	_, errorUpdate := taskService.Database.Exec(query, task.Status, task.Id)
	return errorUpdate
}

func (taskService *TaskService) ListTasks() ([]model.Task, error) {
	var query string = "SELECT * FROM tasks"
	rows, errorGetAll := taskService.Database.Query(query)
	if errorGetAll != nil {
		return nil, errorGetAll
	}

	defer rows.Close()

	var allTasks []model.Task
	for rows.Next() {
		var task model.Task
		errorScan := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Status,
			&task.CreatedAt,
		)

		if errorScan != nil {
			return nil, errorScan
		}

		allTasks = append(allTasks, task)
	}

	return allTasks, nil
}

func (taskService *TaskService) ProccesTasks() {
	for task := range taskService.TaskChannel {
		log.Printf("Processing task %s", task.Title)
		time.Sleep(5 * time.Second)
		task.Status = "completed"
		taskService.UpdateTask(&task)
		log.Printf("Task %s as processed", task.Title)
	}
}
