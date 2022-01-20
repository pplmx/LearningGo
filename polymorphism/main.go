package main

import (
	"LearningGo/polymorphism/demo"
	"github.com/google/uuid"
	"time"
)

func main() {

	var repo demo.Repository
	taskRepository := demo.TaskRepository{}
	repo = taskRepository
	task := demo.Task{
		Id:          uuid.New(),
		CreatedTime: time.Now().UnixMicro(),
		UpdatedTime: time.Now().UnixMicro(),
		DeletedTime: time.Now().UnixMicro(),
	}
	if err := repo.Save(task); err != nil {
		return
	}

}
