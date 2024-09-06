package basic

import (
	"log"
	"my-protobuf/protogen/basic"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteProtoToFile(msg proto.Message, fname string) error {
	b, err := proto.Marshal(msg)
	if err != nil {
		log.Fatalln("Can not marshal message:", err)
	}

	if err := os.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can not write to file:", err)
	}
	return nil
}

func ReadProtoFromFile(fname string, dest proto.Message) {
	in, err := os.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file:", err)
	}

	if err := proto.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal message:", err)
	}
}

func dummyUser() basic.User {
	addr := basic.Address{
		Street:  "Jl. Jend. Sudirman",
		City:    "Jakarta",
		Country: "Indonesia",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  -6.2088,
			Longitude: 106.8456,
		},
	}

	comm := randomCommunicationChannel()

	sr := map[string]uint32{
		"Jl. Jend. Sudirman": 1,
		"Jakarta":            2,
		"Indonesia":          3,
	}

	return basic.User{
		Id:                   1,
		Username:             "user",
		IsActive:             true,
		Password:             []byte("password"),
		Emails:               []string{"superadmin@localhost", "admin@localhost"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
		SkillRating:          sr,
	}
}

func WriteToFileSample() {
	u := dummyUser()
	WriteProtoToFile(&u, "superman_file.bin")
}

func ReadProtoFromFileSample() {
	var u basic.User
	ReadProtoFromFile("superman_file.bin", &u)
	log.Println(&u)
}
