package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicWriteUserContentV1() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/user-content-slug",
		// Title:         "User Content Title",
		// HtmlContent: "<h1>User Content Title</h1><p>This is a user content</p>",
		// AuthorId:      99,
	}
	WriteProtoToFile(&uc, "user_content_v1.bin")
}

func BasicReadUserContentV1() {
	log.Println("Reading user content v1")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v1.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV2() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/user-content-slug-v2",
		// Title:         "User Content Title 2",
		// HtmlContent: "<h1>User Content Title</h1><p>This is a user content</p>",
		// AuthorId:      99,
		// Category: "Category 1",
	}
	WriteProtoToFile(&uc, "user_content_v2.bin")
}

func BasicReadUserContentV2() {
	log.Println("Reading user content v2")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v2.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV3() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/user-content-slug-v3",
		// Title:         "User Content Title 2",
		// HtmlContent: "<h1>User Content Title</h1><p>This is a user content</p>",
		// AuthorId:      99,
		// Category: "Category 1",
		// SubCategory: "Sub Category 1",
	}
	WriteProtoToFile(&uc, "user_content_v3.bin")
}

func BasicReadUserContentV3() {
	log.Println("Reading user content v3")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v3.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}

func BasicWriteUserContentV4() {
	uc := basic.UserContent{
		UserContentId: 1001,
		Slug:          "/user-content-slug-v3",
		Title:         "User Content Title 2",
		// HtmlContent: "<h1>User Content Title</h1><p>This is a user content</p>",
		AuthorId: 99,
		// Rating:   10,
		// Category: "Category 1",
		// SubCategory: "Sub Category 1",
	}
	WriteProtoToFile(&uc, "user_content_v4.bin")
}

func BasicReadUserContentV4() {
	log.Println("Reading user content v4")
	uc := basic.UserContent{}
	ReadProtoFromFile("user_content_v4.bin", &uc)

	log.Println(&uc)

	ucJson, _ := protojson.Marshal(&uc)
	log.Println(string(ucJson))
}
