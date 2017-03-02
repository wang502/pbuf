package structs

import (
    pb "github.com/wang502/pbuf/protobuf"
)

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

var Person_PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var Person_PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

type Person struct {
  Name   string
  Id     int32
  Email  string
  Phones []*pb.Person_PhoneNumber
}

type Person_PhoneNumber struct {
	Number string
	Type   Person_PhoneType
}

type AddressBook struct {
	People []*Person
}
