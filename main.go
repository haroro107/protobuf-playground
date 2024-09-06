package main

import (
	"fmt"
	"log"
	"my-protobuf/basic"
	"time"
)

type logWriter struct{}

func (write logWriter) Write(bytes []byte) (int, error) {
	return fmt.Println(time.Now().Format("15:04:05") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	basic.BasicUser()

	// basic.ProtoToJsonUser()
	// basic.JsontoProtoUser()

	// basic.BasicUserGroup()

	// jobsearch.JobSearchSoftware()
	// jobsearch.JobSearchCandidate()
	// basic.BasicUnmarshalAnyKnown()
	// basic.BasicUnmarshalAnyNotKnown()
	// basic.BasicUnmarshalAnyIs()

	// basic.BasicOneof()
	// basic.WriteToFileSample()
	// basic.ReadProtoFromFileSample()

	// basic.WriteToJSONSample()
	// basic.ReadProtoFromJSONSample()
	// basic.BasicWriteUserContentV1()
	// basic.BasicReadUserContentV1()
	// basic.BasicWriteUserContentV2()
	// basic.BasicReadUserContentV2()
	// basic.BasicWriteUserContentV3()
	// basic.BasicReadUserContentV3()

	// basic.BasicWriteUserContentV4()
	// basic.BasicReadUserContentV4()

	// basic.BasicReadUserPayment()
}
