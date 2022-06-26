package server

import (
	"context"
	protos "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"log"
)

type PipelineStatus struct {
	protos.PipelineStatusServer
}

func NewPipelineStatus() *PipelineStatus {
	return &PipelineStatus{}
}

func (c *PipelineStatus) GetStatus(_ context.Context, req *protos.GetStatusRequest) (*protos.GetStatusResponse, error) {
	log.Println("Handle GetPipelineStatus", "group", req.GetGroup())
	return &protos.GetStatusResponse{
		Branch:  "main",
		Status:  "SUCCESS",
		Commit:  "54fc57g5",
		Author:  "toto",
		Message: "no message",
		Date:    "now",
		Url:     "https://github.com",
	}, nil
}
