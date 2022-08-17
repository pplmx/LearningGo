package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
)

/*
{"taskId":"188388","rootTaskId":"1234","workerId":"a33bdeae-6313-4371-8af8-b77a7b94de75","workerMAC":"D8-BB-C1-06-DE-A8","assignedDate":"2022-08-12 06:03:25.869"}
{"taskId":"188406","rootTaskId":"1234","workerId":"a33bdeae-6313-4371-8af8-b77a7b94de75","workerMAC":"D8-BB-C1-06-DE-A8","assignedDate":"2022-08-12 06:04:18.385"}
{"taskId":"188424","rootTaskId":"1000","workerId":"a33bdeae-6313-4371-8af8-b77a7b94de75","workerMAC":"D8-BB-C1-06-DE-A8","assignedDate":"2022-08-12 06:06:58.042"}
{"taskId":"188440","rootTaskId":"1000","workerId":"a33bdeae-6313-4371-8af8-b77a7b94de75","workerMAC":"D8-BB-C1-06-DE-A8","assignedDate":"2022-08-12 06:07:24.361"}
*/

type Record struct {
	TaskId       string `json:"taskId"`
	RootTaskId   string `json:"rootTaskId"`
	WorkerId     string `json:"workerId"`
	WorkerMAC    string `json:"workerMAC"`
	AssignedDate string `json:"assignedDate"`
}

// search a text by regex in a file
func search(regex string, path string) []Record {
	var result []Record
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error occurred when opening %s.", path))
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(fmt.Sprintf("Error occurred when closing %s.", path))
		}
	}(f)

	// read the file line by line
	scanner := bufio.NewScanner(f)
	re, err := regexp.Compile(regex)
	if err != nil {
		log.Fatalln(fmt.Sprintf("Error occurred when compiling the regex: %s.", regex))
	}
	for scanner.Scan() {
		var r Record
		if re.MatchString(scanner.Text()) {
			err := json.Unmarshal([]byte(scanner.Text()), &r)
			if err != nil {
				log.Fatalln(fmt.Sprintf("Error occurred when unmarshalling the json: %s.", scanner.Text()))
			}
			result = append(result, r)
		}
	}
	return result
}

func searchByTaskId(taskId string, path string) []Record {
	regex := fmt.Sprintf("^{\"taskId\":\"(%s)\",\"rootTaskId\":\"(.*)\",\"workerId\":\"(.*)\",\"workerMAC\":\"(.*)\",\"assignedDate\":\"(.*)\"}$", taskId)
	return search(regex, path)
}

func searchByRootTaskId(rootTaskId string, path string) []Record {
	regex := fmt.Sprintf("^{\"taskId\":\"(.*)\",\"rootTaskId\":\"(%s)\",\"workerId\":\"(.*)\",\"workerMAC\":\"(.*)\",\"assignedDate\":\"(.*)\"}$", rootTaskId)
	return search(regex, path)
}

func searchByMAC(mac string, path string) []Record {
	regex := fmt.Sprintf("^{\"taskId\":\"(.*)\",\"rootTaskId\":\"(.*)\",\"workerId\":\"(.*)\",\"workerMAC\":\"(%s)\",\"assignedDate\":\"(.*)\"}$", mac)
	return search(regex, path)
}

func main() {
	tasks := searchByTaskId("188388", "/var/tmp/re.log")
	fmt.Printf("search tasks by task_id[188388]: %+v\n", tasks)
	subtasks1 := searchByRootTaskId("1234", "/var/tmp/re.log")
	fmt.Printf("search subtasks by root_task_id[1234]: %+v\n", subtasks1)
	subtasks2 := searchByRootTaskId("1000", "/var/tmp/re.log")
	fmt.Printf("search subtasks by root_task_id[1000]: %+v\n", subtasks2)
	tasksOnMAC := searchByMAC("D8-BB-C1-06-DE-A8", "/var/tmp/re.log")
	fmt.Printf("search tasks by MAC[D8-BB-C1-06-DE-A8]: %+v\n", tasksOnMAC)
}
