package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
)

// https://stackoverflow.com/questions/12854125/how-do-i-dump-the-struct-into-the-byte-array-without-reflection/12854659#12854659

type Profile struct {
	Id       int
	Name     string
	password string
}

type Demo struct {
	DemoId      int
	DemoName    string
	UserProfile Profile
}

func EncodeRecursive(encoder *gob.Encoder, v reflect.Value) {
	switch v.Kind() {
	case reflect.Ptr:
		EncodeRecursive(encoder, v)
	case reflect.Interface:
		EncodeRecursive(encoder, v)
	case reflect.Struct:
		EncodeRecursive(encoder, v)
	case reflect.Slice:
		EncodeRecursive(encoder, v)
	case reflect.Map:
		EncodeRecursive(encoder, v)
	default:
		EncodeRecursive(encoder, v)
	}
}

func GobEncode(obj interface{}) ([]byte, error) {
	w := new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	v := reflect.ValueOf(obj)
	EncodeRecursive(encoder, v)
	return w.Bytes(), nil
}

func GobDecode(buf []byte, obj *interface{}) error {
	r := bytes.NewBuffer(buf)
	decoder := gob.NewDecoder(r)

	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		err := decoder.Decode(&obj)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	demo := Demo{
		DemoId:   666,
		DemoName: "Demo Case",
		UserProfile: Profile{
			Id:       999,
			Name:     "Tom",
			password: "Passw0rd",
		},
	}
	fmt.Printf("%+v\n", demo)
	//encode, err := GobEncode(demo)
	//if err != nil {
	//    return
	//}
	//fmt.Println(encode)
}
