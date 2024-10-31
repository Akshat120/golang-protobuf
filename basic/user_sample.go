package basic

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BasicUser() basic.User {
	addr := basic.Address{
		Street:  "L0-12",
		City:    "London",
		Country: "United Kingdom",
		Coordinate: &basic.Address_Coordinate{
			Longitude: 53.1,
			Latitude:  54.1,
		},
	}

	comm := randomCommunicationChannel()

	skillRating := map[string]uint32{
		"c++":    5,
		"java":   7,
		"python": 8,
	}

	u := basic.User{
		Id:                   15,
		Username:             "Clark Kent",
		IsActive:             true,
		Password:             []byte("my-pass"),
		Gender:               basic.Gender_GENDER_MALE,
		Emails:               []string{"abc@gmail.com", "def@gmail.com"},
		Address:              &addr,
		CommunicationChannel: comm,
		SkillRating:          skillRating,
		LastLoginTimestamp:   timestamppb.New(time.Now()),
		BirthDate: &basic.Date{
			Day:   19,
			Month: 7,
			Year:  2000,
		},
		LastKnownLocation: &basic.LatLng{
			Longitude: 54.1343434,
			Latitude:  13.3423423,
		},
	}

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}
	log.Println(string(jsonBytes))

	return u
}

func ProtoToJson() {
	u := basic.User{
		Id:       1,
		Username: "Peter Parker",
		IsActive: false,
		Password: []byte("my-pass"),
		Gender:   basic.Gender_GENDER_MALE,
		Emails:   []string{"abc@gmail.com", "def@gmail.com"},
	}
	log.Println(&u)

	// Marshal the User struct to JSON
	jsonByte, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	log.Println(string(jsonByte))
}

func JsonToProto() {

	json := `{
		"id":2, 
		"username":"Robert Downy Jr", 
		"password":"bXktcGFzcw==",
		"emails":["abc@gmail.com", "def@gmail.com"], 
		"gender":"GENDER_MALE"
	}`

	var u basic.User
	err := protojson.Unmarshal([]byte(json), &u)
	if err != nil {
		log.Fatalf("Failed to unmarshal user: %v", err)
	}

	log.Println(&u)
}

func randomCommunicationChannel() *anypb.Any {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var (
		communicationChannel *anypb.Any
		err                  error
	)
	randomInt := rng.Intn(3)

	if randomInt == 0 {
		socialMedia := &basic.SocialMedia{
			SocialMediaPlatform: "Twitter",
			SocialMediaUsername: "clark_kent",
		}

		communicationChannel, err = anypb.New(socialMedia)
		if err != nil {
			log.Fatalf("Failed to marshal SocialMedia to Any: %v", err)
		}
	} else if randomInt == 1 {

		paperMail := &basic.PaperMail{
			PaperMailAddress: "123 Main St, Metropolis",
		}

		communicationChannel, err = anypb.New(paperMail)
		if err != nil {
			log.Fatalf("Failed to marshal PaperMail to Any: %v", err)
		}
	} else {

		instantMessaging := &basic.InstantMessaging{
			InstantMessagingProduct:  "WhatsApp",
			InstantMessagingUsername: "clarkkent123",
		}
		communicationChannel, err = anypb.New(instantMessaging)
		if err != nil {
			log.Fatalf("Failed to marshal InstantMessaging to Any: %v", err)
		}
	}

	return communicationChannel
}

func BasicMarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "Facebook",
		SocialMediaUsername: "zepurts",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	jsonBytes, err := protojson.Marshal(&sm)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	log.Println(string(jsonBytes))

}

func BasicMarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var (
		unmarshaled protoreflect.ProtoMessage
		err         error
	)

	unmarshaled, err = a.UnmarshalNew()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, err := protojson.Marshal(unmarshaled)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}

	log.Println(string(jsonBytes))

}

func BasicUnmasrshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalf("Failed to unmarshal paperMail: %v", err)
		}

		jsonBytes, err := protojson.Marshal(&pm)
		if err != nil {
			log.Fatalf("Failed to marshal user: %v", err)
		}

		log.Println(string(jsonBytes))
	} else {
		log.Println("Not a PaperMail, but :", a.TypeUrl)
	}

}

func BasicOneOf() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "Facebook",
	// 	SocialMediaUsername: "zepurts",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "Whatsapp",
		InstantMessagingUsername: "akshat",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	u := basic.User{
		Id:       45,
		Username: "my-username",
		IsActive: true,
		Password: []byte("my-password"),
		Address: &basic.Address{
			Street:  "my-street",
			City:    "London",
			Country: "United Kingdom",
			Coordinate: &basic.Address_Coordinate{
				Longitude: 56.234234234,
				Latitude:  52.6764543534,
			},
		},
		Emails:                []string{"abc@gmai.com", "def@gmail.com"},
		Gender:                basic.Gender_GENDER_MALE,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		log.Fatalf("Failed to marshal user: %v", err)
	}
	log.Println(string(jsonBytes))

}

/*
Write protobuf to disk
*/
func WriteProtoToDisk() {
	u := BasicUser()

	bytes, err := proto.Marshal(&u)
	if err != nil {
		fmt.Println("Error:", err)
	}

	ioutil.WriteFile("user_proto.bin", bytes, 0777)

}

/*
Read protobuf from disk and convert into json
*/
func ReadProtoFromDisk() {
	bytes, err := ioutil.ReadFile("user_proto.bin")
	if err != nil {
		fmt.Println("Error:", err)
	}

	user := basic.User{}

	err = proto.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	jsonBytes, err := protojson.Marshal(&user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	log.Println(string(jsonBytes))
}

/*
Write json to disk
*/
func WriteJsonToDisk() {
	u := BasicUser()

	jsonBytes, err := protojson.Marshal(&u)
	if err != nil {
		fmt.Println("Error:", err)
	}

	ioutil.WriteFile("user_json.bin", jsonBytes, 0777)
}

/*
Reads json from disk and convert into protobuf
*/
func ReadJsonFromDisk() {
	bytes, err := ioutil.ReadFile("user_json.bin")
	if err != nil {
		fmt.Println("Error:", err)
	}

	user := basic.User{}

	err = protojson.Unmarshal(bytes, &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	log.Println(&user)
}
