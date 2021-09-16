package main

import (
    "bufio"
    "crypto/rand"
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "os"
    "strconv"
    "strings"
    "time"
)

var count = 0
var ServerCertFile = "./certificate/localhost.crt.pem"
var ServerKeyFile = "./certificate/localhost.key.pem"

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

    // load certificate
    crt, err := tls.LoadX509KeyPair(ServerCertFile, ServerKeyFile)
    if err != nil {
        log.Fatalln(err.Error())
        return
    }
    tlsConfig := &tls.Config{}
    tlsConfig.Certificates = []tls.Certificate{crt}
    tlsConfig.Time = time.Now
    tlsConfig.Rand = rand.Reader

    l, err := tls.Listen("tcp", PORT, tlsConfig)
    if err != nil {
        log.Fatalln(err.Error())
        return
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println(err.Error())
            continue
        }
        go handleConnection(conn)
        count++
    }
}
