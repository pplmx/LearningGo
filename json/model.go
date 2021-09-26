package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
)

type Demo struct {
	Id      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	Batters struct {
		Batter []struct {
			Id   string `json:"id"`
			Type string `json:"type"`
		} `json:"batter"`
	} `json:"batters"`
	Topping []struct {
		Id   string `json:"id"`
		Type string `json:"type"`
	} `json:"topping"`
}

func main() {
    content, err := os.ReadFile("data1.json")
    if err != nil {
        log.Fatal("Error when opening file: ", err)
    }

    var demo Demo
    err = json.Unmarshal(content, &demo)
    if err != nil {
        log.Fatal("Error during Unmarshal(): ", err)
    }
    bytesVar, err := json.Marshal(demo)
    fmt.Println(demo)
    fmt.Println(string(bytesVar))

}
