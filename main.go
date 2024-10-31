package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
	// basic.BasicHello()
	basic.BasicUser()
	// basic.ProtoToJson()
	// basic.JsonToProto()
	// basic.BasicUserGroup()
	// basic.BasicJobApplication()
	// basic.BasicSoftwareApplication()
	// basic.BasicMarshalAnyKnown()
	// basic.BasicMarshalAnyNotKnown()
	// basic.BasicUnmasrshalAnyIs()
	// basic.BasicOneOf()
	// basic.WriteProtoToDisk()
	// basic.ReadProtoFromDisk()
	// basic.WriteJsonToDisk()
	// basic.ReadJsonFromDisk()
	// basic.WriteToFileSample()
	// basic.ReadFromFileSample()
}
