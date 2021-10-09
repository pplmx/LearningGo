package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "reflect"
)

// https://stackoverflow.com/questions/12854125/how-do-i-dump-the-struct-into-the-byte-array-without-reflection/12854659#12854659

type Profile struct {
    Id int
    Name string
    password string
}

type Demo struct {
    demoId int
    demoName string
    userProfile Profile
}

func GobEncode(obj interface{}) ([]byte, error) {
    w := new(bytes.Buffer)
    encoder := gob.NewEncoder(w)

    v := reflect.ValueOf(obj)
    fmt.Println(v.NumField())
    for i := 0; i < v.NumField(); i++ {
        err := encoder.Encode(v.Field(i).Interface())
        if err != nil {
            return nil, err
        }
    }
    return w.Bytes(), nil
}

//func GobDecode(buf []byte, obj *interface{}) (*interface{}, error) {
//    r := bytes.NewBuffer(buf)
//    decoder := gob.NewDecoder(r)
//
//    v := reflect.ValueOf(obj)
//    for i := 0; i < v.NumField(); i++ {
//        err := decoder.Decode(v.Field(i).Interface())
//        if err != nil {
//            return nil, err
//        }
//    }
//    return decoder.Decode(&d.name)
//}

func main() {
    demo := Demo{
        demoId:      666,
        demoName:    "Demo Case",
        userProfile: Profile{
            Id: 999,
            Name: "Tom",
            password: "Passw0rd",
        },
    }
    encode, err := GobEncode(demo)
    if err != nil {
        return
    }
    fmt.Println(encode)
}
