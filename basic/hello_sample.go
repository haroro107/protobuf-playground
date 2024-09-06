package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Hello, World!",
	}
	log.Println(&h)
}
