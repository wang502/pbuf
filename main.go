package main

import (
    "bytes"
    "fmt"
    "github.com/golang/protobuf/proto"
    pb "github.com/wang502/pbuf/protobuf"
    "github.com/wang502/pbuf/structs"
)

func main()  {
    pbPerson := &pb.Person{
        Name: "Xiaohui Wang",
        Id: 10,
        Email: "example@pinterest.com",
        Phones: []*pb.Person_PhoneNumber{
            {Number: "555-4321", Type: pb.Person_HOME},
        },
    }

    var encodedBuf bytes.Buffer
    p, err := proto.Marshal(pbPerson)
    if err != nil {
        fmt.Println(err)
        return
    }
    encodedBuf.Write(p)

    decodedPerson := new(pb.Person)
    if err := proto.Unmarshal(encodedBuf.Bytes(), decodedPerson); err!= nil {
        fmt.Println(err)
        return
    }

    person := new(structs.Person)
    person.Name = decodedPerson.GetName()
    person.Id = decodedPerson.GetId()
    person.Email = decodedPerson.GetEmail()
    person.Phones = decodedPerson.GetPhones()

    fmt.Println(person)

    // ------------------------------------------------------------

    person = &structs.Person{
      Name: "Xiaohui Wang",
      Id: 10,
      Email: "example@pinterest.com",
      Phones: []*pb.Person_PhoneNumber{
          {Number: "555-4321", Type: pb.Person_HOME},
      },
    }
    var w bytes.Buffer
    _, err = person.Encode(&w)
    if err != nil {
        fmt.Println(err)
        return
    }

    emptyPerson := &structs.Person{}
    _, err = emptyPerson.Decode(&w)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(emptyPerson)
}
