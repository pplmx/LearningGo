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

const (
	HeaderLength = 2 + 8 + 2 + 8
)

// Header length is a constant: 20B
type Header struct {
	Version            [2]byte // protocol version
	Timestamp          int64
	CommonStreamLength int16 // Maximum 64KB
	BigStreamLength    int64 // Maximum 16EB
}

type Body struct {
	CommonStream []byte
	BigStream    []byte
}

type Pck struct {
	head Header
	body Body
}

func (p *Pck) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.head.Version)
	err = binary.Write(writer, binary.BigEndian, &p.head.Timestamp)
	err = binary.Write(writer, binary.BigEndian, &p.head.CommonStreamLength)
	err = binary.Write(writer, binary.BigEndian, &p.head.BigStreamLength)

	err = binary.Write(writer, binary.BigEndian, &p.body.CommonStream)
	err = binary.Write(writer, binary.BigEndian, &p.body.BigStream)
	return err
}

func (p *Pck) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.head.Version)
	err = binary.Read(reader, binary.BigEndian, &p.head.Timestamp)
	err = binary.Read(reader, binary.BigEndian, &p.head.CommonStreamLength)
	err = binary.Read(reader, binary.BigEndian, &p.head.BigStreamLength)

	p.body.CommonStream = make([]byte, p.head.CommonStreamLength)
	err = binary.Read(reader, binary.BigEndian, &p.body.CommonStream)
	p.body.BigStream = make([]byte, p.head.BigStreamLength)
	err = binary.Read(reader, binary.BigEndian, &p.body.BigStream)
	return err
}

func (p *Pck) String() string {
	return fmt.Sprintf("head:%+v body.common:%s, body.big:%s",
		p.head,
		p.body.CommonStream,
		p.body.BigStream,
	)
}

func main() {
	pck := &Pck{
		head: Header{
			Version:   [2]byte{'v', '1'},
			Timestamp: time.Now().Unix(),
		},
		body: Body{
			CommonStream: []byte(("现在时间是:" + time.Now().Format("2006-01-02 15:04:05"))),
			BigStream:    []byte(("Big Stream:" + time.Now().Format("2006-01-02 15:04:05"))),
		},
	}
	pck.head.CommonStreamLength = int16(len(pck.body.CommonStream))
	pck.head.BigStreamLength = int64(len(pck.body.BigStream))

	buf := new(bytes.Buffer)
	// write four times to simulate the sticking packages
	err := pck.Pack(buf)
	if err != nil {
		return
	}
	pck.Pack(buf)
	pck.Pack(buf)
	pck.Pack(buf)
	// scanner
	scanner := bufio.NewScanner(buf)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'v' {
			if len(data) > HeaderLength { // It means data segment is not empty.
				commonStreamLength := int16(0)
				bigStreamLength := int64(0)
				binary.Read(bytes.NewReader(data[10:12]), binary.BigEndian, &commonStreamLength)
				binary.Read(bytes.NewReader(data[12:20]), binary.BigEndian, &bigStreamLength)
				// TODO Potential bug: the total sum can be over the range of int64
				packageLength := int64(HeaderLength) + int64(commonStreamLength) + bigStreamLength
				if packageLength <= int64(len(data)) {
					// TODO int64 convert to int, which can loss the data
					return int(packageLength), data[:packageLength], nil
				}
			}
		}
		return
	})
	for scanner.Scan() {
		scannedPack := new(Pck)
		scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		log.Println(scannedPack)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Invalid packet.")
	}
}
