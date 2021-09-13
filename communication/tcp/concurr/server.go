package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn) {
    fmt.Println("Connection from ", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(counter))
	}
	defer func(c net.Conn) {
        err := c.Close()
        if err != nil {
            fmt.Println("Failed to Close Connection from ", c.RemoteAddr().String())
        }
    }(c)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}
