package jobsearch

import (
	"log"
	"my-protobuf/protogen/basic"
	"my-protobuf/protogen/dummy"
	"my-protobuf/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "JobSearch",
			Platforms: []string{"Windows", "Linux", "MacOS"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software:", string(jsonBytes))
}

func JobSearchCandidate() {
	js := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId:       999,
			ApplicationFullName: "Aldi",
			Phone:               "123-456-7890",
			Email:               "aldi@gmail.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Candidate:", string(jsonBytes))
}
