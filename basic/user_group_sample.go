package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {

	us := []*basic.User{
		{
			Id:       1,
			Username: "Peter Parker",
			IsActive: false,
			Password: []byte("my-pass"),
			Gender:   basic.Gender_GENDER_MALE,
			Emails:   []string{"abc@gmail.com", "def@gmail.com"},
		},
		{
			Id:       2,
			Username: "Mary Jane",
			IsActive: false,
			Password: []byte("my-pass"),
			Gender:   basic.Gender_GENDER_FEMALE,
			Emails:   []string{"abc@gmail.com", "def@gmail.com"},
		},
	}

	ug := basic.UserGroup{
		GroupId:     1,
		GroupName:   "Joker",
		Roles:       []string{"R1", "R2"},
		Users:       us,
		Description: "MY-DESCRIPTION",
	}

	json, err := protojson.Marshal(&ug)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	log.Println(string(json))
}
