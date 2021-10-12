package main

import (
    "LearningGo/communication/encryption/domain"
    "crypto/rand"
    "crypto/tls"
    "encoding/json"
    "log"
    "net"
    "time"
)

func main() {
    serverCertFile := "./certificate/localhost.crt.pem"
    serverKeyFile := "./certificate/localhost.key.pem"

    startServer("12345", serverKeyFile, serverCertFile)
}

func startServer(port string, keyFile string, certFile string) {
    // load certificate
    crt, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        log.Fatalln(err)
    }
    tlsConfig := &tls.Config{}
    tlsConfig.Certificates = []tls.Certificate{crt}
    tlsConfig.Time = time.Now
    tlsConfig.Rand = rand.Reader

    listener, err := tls.Listen("tcp", ":"+port, tlsConfig)
    if err != nil {
        log.Fatalln(err)
    }
    defer func(l net.Listener) {
        err := l.Close()
        if err != nil {
            log.Fatalln(err)
        }
    }(listener)
    log.Println("Listening on", listener.Addr().String())

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Println(err)
            continue
        }
        // create a coroutine(goroutine)
        go handleConnection(conn)
    }
}

func handleConnection(c net.Conn) {
    log.Println("Connection from", c.RemoteAddr().String())
    defer func(c net.Conn) {
        err := c.Close()
        if err != nil {
            log.Fatalln("Failed to Close Connection from", c.RemoteAddr().String())
        }
    }(c)

    for {
        // parse the request
        var req domain.Request
        err := json.NewDecoder(c).Decode(&req)
        if err != nil {
            log.Println("Client exited due to", err)
            return
        }
        log.Printf("%+v\n", req)
        _, err = c.Write([]byte("Received\n"))
        if err != nil {
            log.Fatalln(err)
        }
    }
}
