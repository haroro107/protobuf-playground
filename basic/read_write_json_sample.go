package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) error {
	b, err := protojson.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message:", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can not write to file:", err)
	}
	return nil
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file:", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal message:", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadProtoFromJSONSample() {
	var u basic.User
	ReadProtoFromJSON("superman_file.json", &u)
	log.Println(&u)
}
