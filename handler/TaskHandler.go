package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andreparelho/task-manager-golang/database"
	"github.com/andreparelho/task-manager-golang/model"
	"github.com/andreparelho/task-manager-golang/service"
)

func HandleCreateTask(responseWriter http.ResponseWriter, request *http.Request) {
	database := database.NewDatabaseConnection()
	defer database.Close()

	taskService := service.TaskConstructor(database, make(chan model.Task))
	var task model.Task

	errorDecode := json.NewDecoder(request.Body).Decode(&task)
	if errorDecode != nil {
		http.Error(responseWriter, errorDecode.Error(), http.StatusBadRequest)
		return
	}

	task.Status = "pending"
	task.CreatedAt = time.Now()

	errorAddTask := taskService.AddTask(task)
	if errorAddTask != nil {
		http.Error(responseWriter, errorAddTask.Error(), http.StatusInternalServerError)
		return
	}

	taskService.TaskChannel <- task

	responseWriter.WriteHeader(http.StatusCreated)
}

func HandleListTask(responseWriter http.ResponseWriter, request *http.Request) {
	database := database.NewDatabaseConnection()
	defer database.Close()

	taskService := service.TaskConstructor(database, nil)

	tasks, errorListTasks := taskService.ListTasks()
	if errorListTasks != nil {
		http.Error(responseWriter, errorListTasks.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(tasks)
}
