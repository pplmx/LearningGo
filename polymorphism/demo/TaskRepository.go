package demo

import "fmt"

type TaskRepository struct {
}

func (r TaskRepository) Save(obj interface{}) error {
	var err error
	if task, ok := obj.(Task); ok {
		fmt.Println(task)
		return err
	}
	return err
}

func (r TaskRepository) FindAll() ([]interface{}, error) {
	var err error
	var tasks []Task
	var ret []interface{}
	for _, task := range tasks {
		ret = append(ret, task)
	}
	return ret, err
}

func (r TaskRepository) FindById(id interface{}) (interface{}, error) {
	var err error
	var task Task
	if _, ok := id.(int); ok {
		return task, err
	}
	return nil, err
}

func (r TaskRepository) DeleteById(id interface{}) error {
	var err error
	if _, ok := id.(int); ok {
		return nil
	}
	return err
}
