package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       1,
		Username: "batman",
		IsActive: true,
		Password: []byte("password"),
		Emails:   []string{"batman@localhost", "brucewayne@localhost"},
		Gender:   basic.Gender_GENDER_MALE,
	}

	nightwing := basic.User{
		Id:       2,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("password"),
		// Emails:   []string{"nightwing@localhost", "dickgrayson@localhost"},
		Gender: basic.Gender_GENDER_UNSPECIFIED,
	}

	batFamily := basic.UserGroup{
		GroupId:     111,
		GroupName:   "Bat Family",
		Users:       []*basic.User{&batman, &nightwing},
		Description: "A group of vigilante crime fighters appearing in American comic books published by DC Comics.",
	}

	jsonByets, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonByets))
}
