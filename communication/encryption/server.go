package main

import (
    "bufio"
    "crypto/rand"
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "io/ioutil"
    "log"
    "net"
    "os"
    "strconv"
    "strings"
    "time"
)

var count = 0

func main() {

    arguments := os.Args
    if len(arguments) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }
    PORT := ":" + arguments[1]

    serverCertFile := "./certificate/localhost.crt.pem"
    serverKeyFile := "./certificate/localhost.key.pem"

    caCert, err := ioutil.ReadFile(serverCertFile)
    if err != nil {
        log.Fatalln(err.Error())
        return
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // load certificate
    crt, err := tls.LoadX509KeyPair(serverCertFile, serverKeyFile)
    if err != nil {
        log.Fatalln(err.Error())
        return
    }
    tlsConfig := &tls.Config{}
    tlsConfig.Certificates = []tls.Certificate{crt}
    tlsConfig.Time = time.Now
    tlsConfig.Rand = rand.Reader
    //tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
    //tlsConfig.ClientCAs = caCertPool

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
