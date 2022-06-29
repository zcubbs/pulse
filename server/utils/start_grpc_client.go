package utils

import (
	"crypto/tls"
	"crypto/x509"
	protos "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

var PipelinesGrpcClient protos.PipelineStatusClient

func StartGrpcClient() func() {
	// Load our TLS certificate and use grpc/credentials to create new transport credentials
	c := credentials.NewTLS(loadTLSCfg())
	conn, err := grpc.Dial("localhost:9092", grpc.WithTransportCredentials(c))
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

// loadTLSCfg will load a certificate and create a tls config
func loadTLSCfg() *tls.Config {
	b, _ := ioutil.ReadFile("../_cert/server.crt")
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		log.Fatal("credentials: failed to append certificates")
	}
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            cp,
	}
	return config
}
