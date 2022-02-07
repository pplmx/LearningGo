package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type PO struct {
	_           [0]func()
	_           struct{}
	ID          string `redis:"id"`
	CreatedTime int64  `redis:"created_time"`
	UpdatedTime int64  `redis:"updated_time"`
	DeletedTime int64  `redis:"deleted_time"`
}

type Task struct {
	PO
	TaskId   int     `redis:"task_id"` // from the dragon
	Error    int     `redis:"error"`
	Status   int     `redis:"status"`
	IsAsync  bool    `redis:"is_async"`
	Results  string  `redis:"results"`
	Priority float64 `redis:"priority"`
	Progress float64 `redis:"progress"`
}

func FindByID(ctx context.Context, rdb *redis.Client, id interface{}) (interface{}, error) {
	var err error
	var task Task
	if id, ok := id.(string); ok {
		hashResultSet := rdb.HGetAll(ctx, "task_hash_"+id)
		err = hashResultSet.Scan(&task.PO)
		err = hashResultSet.Scan(&task)
	} else {
		err = errors.New("failed to convert id to string type")
	}
	if err != nil {
		return nil, err
	}
	return task, err
}

func FindAll(ctx context.Context, rdb *redis.Client) ([]interface{}, error) {
	var err error
	var retArr []interface{}
	iter := rdb.ScanType(ctx, 0, "task_hash_*", 100, "hash").Iterator()
	var task Task
	for iter.Next(ctx) {
		hashResultSet := rdb.HGetAll(ctx, iter.Val())
		err = hashResultSet.Scan(&task.PO)
		err = hashResultSet.Scan(&task)
		retArr = append(retArr, task)
	}
	return retArr, err
}

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	task, err := FindByID(ctx, rdb, "ee1c1e1e-d928-46a9-9948-68f50a0bef6d")
	if err != nil {
		return
	}
	fmt.Println(fmt.Sprintf("task: %+v", task))

	all, err := FindAll(ctx, rdb)
	if err != nil {
		return
	}
	for _, v := range all {
		fmt.Println(fmt.Sprintf("task: %+v", v))
	}

}
