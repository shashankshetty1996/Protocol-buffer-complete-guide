package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	complexpb "github.com/shashankshetty1996/protobuf-example-go/src/complex"
	enumpb "github.com/shashankshetty1996/protobuf-example-go/src/enum_example"
	simplepb "github.com/shashankshetty1996/protobuf-example-go/src/simple"
)

func main() {
	// With simple proto file
	sm := doSimple()
	ReadAndWriteDemo(sm)
	jsonDemo(sm)

	// With enum type
	doEnum()

	// With complex proto
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "first message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "third message",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}
	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJson(sm)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJson(smAsString, sm2)
	fmt.Println("Successfully created proto struct: ", sm2)
}

func toJson(pb proto.Message) string {
	marshaller := jsonpb.Marshaler{}
	out, err := marshaller.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return ""
	}
	return out
}

func fromJson(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't Unmarshal the JSON into the pb struct", err)
	}
}

// proto.Message or *simplepb.SimpleMessage - both works
func ReadAndWriteDemo(sm *simplepb.SimpleMessage) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println("Reading from file content ", sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to the file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Something went wrong while reading", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't put bytes into the protocol buffers struct", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println(sm)
	fmt.Printf("The ID is %v\n", sm.GetId())

	return &sm
}
