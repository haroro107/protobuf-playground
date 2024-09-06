package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"
	"time"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BasicUser() {
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

	u := basic.User{
		Id:                   1,
		Username:             "user",
		IsActive:             true,
		Password:             []byte("password"),
		Emails:               []string{"superadmin@localhost", "admin@localhost"},
		Gender:               basic.Gender_GENDER_MALE,
		Address:              &addr,
		CommunicationChannel: &comm,
		SkillRating:          sr,
		LastLoginTimestamp:   timestamppb.Now(),
		BirthDate:            &date.Date{Year: 2000, Month: 5, Day: 27},
		LastKnownLocation: &latlng.LatLng{
			Latitude:  -6.29847717,
			Longitude: 106.8290577,
		},
	}
	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
	log.Println(&u)
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       1,
		Username: "Woman",
		IsActive: true,
		Password: []byte("password"),
		Emails:   []string{"superadmin@localhost", "admin@localhost"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsontoProtoUser() {
	json := `{"id":1,"username":"Woman","is_active":true,"password":"cGFzc3dvcmQ=","emails":["superadmin@localhost","admin@localhost"],"gender":"GENDER_FEMALE"}`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(&p)
}

func randomCommunicationChannel() anypb.Any {
	rand.Seed(time.Now().UnixNano())

	paperMail := basic.PaperMail{
		PaperMailAddress: "some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "tester.name",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "what's up",
		InstantMessagingUsername: "tester.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-platform",
		SocialMediaUsername: "my-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshaled as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Message is not PaperMail, but :", a.TypeUrl)
	}

}

func BasicOneof() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "my-platform",
	// 	SocialMediaUsername: "my-username",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "my-instant",
		InstantMessagingUsername: "my-king",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	u := basic.User{
		Id:                    1,
		Username:              "user",
		IsActive:              true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
