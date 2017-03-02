package structs

import (
    "io"
    "io/ioutil"
    "github.com/golang/protobuf/proto"
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

// Encode the Person to a buffer.
// Returns number of bytes written and any error occurs
func (p *Person) Encode(w io.Writer) (int, error) {
    pb := &pb.Person{
        Name    : p.Name,
        Id      : p.Id,
        Email   : p.Email,
        Phones  : p.Phones,
    }
    data, err := proto.Marshal(pb)
    if err != nil {
        return -1, err
    }

    return w.Write(data)
}

// Decode the Person from a buffer.
// Returns the number of bytes read and any error if occurs
func (p *Person) Decode(r io.Reader) (int, error) {
    data, err := ioutil.ReadAll(r)
    if err != nil {
        return -1, err
    }

    pb := &pb.Person{}
    if err := proto.Unmarshal(data, pb); err != nil {
        return -1, err
    }

    p.Name = pb.GetName()
    p.Id = pb.GetId()
    p.Email = pb.GetEmail()
    p.Phones = pb.GetPhones()

    return len(data), nil
}

type Person_PhoneNumber struct {
	Number string
	Type   Person_PhoneType
}

type AddressBook struct {
	People []*Person
}
