package main

import (
    "bufio"
    "bytes"
    "encoding/binary"
    "fmt"
    "io"
    "log"
    "time"
)

type Package struct {
    Version   [2]byte // protocol version
    Length    int16   // Data length
    Timestamp int64
    Data      []byte
}

func (p *Package) Pack(writer io.Writer) error {
    var err error
    err = binary.Write(writer, binary.BigEndian, &p.Version)
    err = binary.Write(writer, binary.BigEndian, &p.Length)
    err = binary.Write(writer, binary.BigEndian, &p.Timestamp)
    err = binary.Write(writer, binary.BigEndian, &p.Data)
    return err
}
func (p *Package) Unpack(reader io.Reader) error {
    var err error
    err = binary.Read(reader, binary.BigEndian, &p.Version)
    err = binary.Read(reader, binary.BigEndian, &p.Length)
    err = binary.Read(reader, binary.BigEndian, &p.Timestamp)
    p.Data = make([]byte, p.Length)
    err = binary.Read(reader, binary.BigEndian, &p.Data)
    return err
}

func (p *Package) String() string {
    return fmt.Sprintf("version:%s length:%d timestamp:%d data:%s",
        p.Version,
        p.Length,
        p.Timestamp,
        p.Data,
    )
}

func main() {
    pack := &Package{
        Version:   [2]byte{'v', '1'},
        Timestamp: time.Now().Unix(),
        Data:      []byte(("现在时间是:" + time.Now().Format("2006-01-02 15:04:05"))),
    }
    pack.Length = int16(len(pack.Data))

    buf := new(bytes.Buffer)
    // 写入四次，模拟TCP粘包效果
    err := pack.Pack(buf)
    if err != nil {
        return
    }
    pack.Pack(buf)
    pack.Pack(buf)
    pack.Pack(buf)
    // scanner
    scanner := bufio.NewScanner(buf)
    scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        if !atEOF && data[0] == 'v' {
            if len(data) > 4 { // It means data segment is not empty.
                length := int16(0)
                binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
                if int(length)+12 <= len(data) {
                    return int(length) + 12, data[:int(length)+12], nil
                }
            }
        }
        return
    })
    for scanner.Scan() {
        scannedPack := new(Package)
        scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
        log.Println(scannedPack)
    }
    if err := scanner.Err(); err != nil {
        log.Fatal("无效数据包")
    }
}
