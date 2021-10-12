package main

import (
    "LearningGo/communication/encryption/domain"
    "bufio"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    "fmt"
    "log"
    "net"
    "os"
    "strings"
)

func main() {
    serverCertFile := "./certificate/localhost.crt.pem"
    conn := connect("localhost:12345", serverCertFile)
    sendRequest(conn)
}

// connect serverAddr e.g. 127.0.0.1:8080 or [::1]:9999, or localhost:8888
func connect(serverAddr string, serverCertFile string) net.Conn {
    caCert, err := os.ReadFile(serverCertFile)
    if err != nil {
        log.Fatalln(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)
    tlsConfig := &tls.Config{
        RootCAs: caCertPool,
    }

    c, err := tls.Dial("tcp", serverAddr, tlsConfig)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Connected with the server", c.RemoteAddr().String())
    return c
}

func sendRequest(c net.Conn) {
    req := domain.Request{
        ReqType: domain.TYPE1,
        ReqData: domain.ReqData{
            Id:   "1",
            Name: "cc",
        },
        Signal: 1,
    }
    err := json.NewEncoder(c).Encode(req)
    if err != nil {
        log.Fatalln(err)
    }
}

func chat(c net.Conn) {
    for {
        // send your message to the server
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        text, _ := reader.ReadString('\n')
        fmt.Fprintf(c, text+"\n")

        // receive the message from the server
        message, _ := bufio.NewReader(c).ReadString('\n')
        fmt.Print("->: " + message)
        if strings.TrimSpace(text) == "STOP" {
            fmt.Println("TCP client exiting...")
            return
        }
    }
}
