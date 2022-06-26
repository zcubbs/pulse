package main

import (
	pb "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"github.com/zcubbs/pulse/pipelines/server"
	"github.com/zcubbs/pulse/pipelines/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GrpcServer struct {
	pb.PipelineStatusServer
}

var addr = ":9092"

func main() {

	// Init Database
	utils.ConnectToPostgresDB()

	// Load yaml config
	utils.LoadYamlConfig()

	// Init Rabbitmq connection
	rabbitmq, channel := utils.ConnectRabbitmq()
	defer channel.Close()
	defer rabbitmq.Close()

	// Setup Git Webhooks
	utils.SetupGitlabWebhook()

	// Launch event worker routine
	utils.LaunchEventWorker()

	// Create a new grpc server
	s := grpc.NewServer()
	ps := server.NewPipelineStatus()

	// Register the server
	pb.RegisterPipelineStatusServer(s, ps)

	// TODO: add a flag-switch to enable/disable this feature
	reflection.Register(s)

	// Start the server
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("listening on %s. ", addr)
	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("Successfully started gRPC server")

}
