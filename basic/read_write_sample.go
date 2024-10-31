package basic

import (
	"fmt"
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/proto"
)

func writeProtoToFile(msg proto.Message, fileName string) {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ioutil.WriteFile(fileName, bytes, 0777)
}

func readProtoFromFile(fileName string, dest proto.Message) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err = proto.Unmarshal(bytes, dest); err != nil {
		fmt.Println("Error:", err)
		return
	}

	log.Println(dest)
}

func dummyUserContent() basic.UserContent {
	return basic.UserContent{
		UserContentId: 1,
		Slug:          "my-slug",
		// Title:         "my-title",
		HtmlContent: "my-html-content",
		// AuthorId:      452,
		// Category: "my-category",
		// SubCategory: "my-sub-category",
	}
}

func WriteToFileSample() {
	uc := dummyUserContent()

	writeProtoToFile(&uc, "user_content_v3.bin")
}

func ReadFromFileSample() {
	readProtoFromFile("user_content_v3.bin", &basic.UserContent{})
}
