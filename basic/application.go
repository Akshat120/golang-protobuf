package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/job"
	"my-protobuf/protogen/software"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicJobApplication() {
	ja := basic.JobApplication{
		Application: &job.Application{
			ApplicationId:     12,
			ApplicantFullName: "Akshat Dhiman",
			Phone:             "123-456-7890",
			Email:             "abc@gmail.com",
		},
		JobId: 33,
	}

	jsonBytes, err := protojson.Marshal(&ja)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}
	log.Println(string(jsonBytes))
}

func BasicSoftwareApplication() {
	sa := basic.SoftwareApplication{
		Application: &software.Application{
			Version:   "1.1.1",
			Name:      "MY-APP",
			Platforms: []string{"iOS", "Android", "MacOS"},
		},
		SoftwareId: 432,
	}

	jsonBytes, err := protojson.Marshal(&sa)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}
	log.Println(string(jsonBytes))
}
