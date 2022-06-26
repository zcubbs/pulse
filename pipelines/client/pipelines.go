package main

import (
	"context"
	protos "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"log"
)

func doGetStatus(c protos.PipelineStatusClient) {
	log.Println("Calling GetPipelineStatus")
	resp, err := c.GetStatus(context.Background(), &protos.GetStatusRequest{Group: "main"})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Printf("Response: %v", resp)
}
