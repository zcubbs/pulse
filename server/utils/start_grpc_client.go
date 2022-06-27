package utils

import (
	protos "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var PipelinesGrpcClient protos.PipelineStatusClient

func StartGrpcClient() func() {
	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to retrieve entries: %s", err)
	}
	log.Println("Connected to grpc server")
	PipelinesGrpcClient = protos.NewPipelineStatusClient(conn)

	// to allow caller to defer conn.Close()
	return func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Failed to close connection: %s", err)
		}
	}
}
