package main

import (
	"github.com/google/uuid"
	"github.com/pplmx/LearningGo/polymorphism/demo"
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
