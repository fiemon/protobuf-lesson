package main

import (
	"fmt"
	// "log"

	"google.golang.org/protobuf/encoding/protojson"

	"protobuf-lesson/pb"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "090-1234-5678"},
		Project:     map[string]*pb.Company_Project{"Projectx": &pb.Company_Project{}},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Day:   1,
		},
	}

	// binDate, err := proto.Marshal(employee)
	// if err != nil {
	// 	log.Fatalln("Can't serialize", err)
	// }

	// if err := os.WriteFile("test.bin", binDate, 0666); err != nil {
	// 	log.Fatalln("Can't write", err)
	// }

	// in, err := os.ReadFile("test.bin")
	// if err != nil {
	// 	log.Fatalln("Can't read file", err)
	// }

	// readEmployee := &pb.Employee{}

	// err = proto.Unmarshal(in, readEmployee)
	// if err != nil {
	// 	log.Fatalln("Can't deserialize", err)
	// }

	// fmt.Println(readEmployee)

	// pb to json
	m := protojson.MarshalOptions{}
	out := m.Format(employee)
	// if err != nil {
	// 	log.Fatalln("Can't marshal to json", err)
	// }

	fmt.Println(out)

}
