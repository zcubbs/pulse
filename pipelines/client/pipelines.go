package main

import (
	"context"
	protos "github.com/zcubbs/pulse/pipelines/proto"
	"io"
	"log"
)

func doGetStatus(c protos.PipelineStatusClient) {
	log.Println("Calling GetPipelineStatus")
	stream, err := c.GetStatus(context.Background(), &protos.GetStatusRequest{Group: "main"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error while reading the stream:", err)
		}

		log.Println("Received:", msg)
	}
}
