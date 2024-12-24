package main

import (
	"log"
	"net/http"

	"github.com/andreparelho/task-manager-golang/database"
	"github.com/andreparelho/task-manager-golang/handler"
	"github.com/andreparelho/task-manager-golang/model"
	"github.com/andreparelho/task-manager-golang/service"
)

func main() {
	database := database.NewDatabaseConnection()
	channel := make(chan model.Task)

	taskService := service.TaskConstructor(database, channel)
	go taskService.ProccesTasks()

	http.HandleFunc("POST /task", handler.HandleCreateTask)
	http.HandleFunc("GET /task", handler.HandleCreateTask)

	log.Println("Application is running on port 8080")

	http.ListenAndServe(":8081", nil)
}
