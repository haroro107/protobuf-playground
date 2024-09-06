package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicReadUserPayment() {
	log.Println("Reading user payment")
	up := basic.UserPayment{}

	ReadProtoFromFile("user_content_v1.bin", &up)

	log.Println(&up)

	upJson, _ := protojson.Marshal(&up)
	log.Println(string(upJson))

}
