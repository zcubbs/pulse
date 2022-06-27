package server

import (
	"github.com/zcubbs/pulse/pipelines/models"
	protos "github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"github.com/zcubbs/pulse/pipelines/queries"
	"log"
	"sort"
)

type PipelineStatus struct {
	protos.PipelineStatusServer
}

func NewPipelineStatus() *PipelineStatus {
	return &PipelineStatus{}
}

func (c *PipelineStatus) GetStatus(req *protos.GetStatusRequest, stream protos.PipelineStatus_GetStatusServer) error {
	log.Println("Handle GetPipelineStatus", "group", req.GetGroup())
	//return &protos.GetStatusResponse{
	//	Branch:  "main",
	//	Status:  "SUCCESS",
	//	Commit:  "54fc57g5",
	//	Author:  "toto",
	//	Message: "no message",
	//	Date:    "now",
	//	Url:     "https://github.com",
	//}, nil

	watchEvents := queries.GetProjectLatestPipelineStatusEntries()
	sort.Sort(models.ByEventDate(watchEvents))
	for _, watchEvent := range watchEvents {
		err := stream.Send(&protos.GetStatusResponse{
			Branch:   watchEvent.Branch,
			Status:   watchEvent.Status,
			Commit:   watchEvent.Commit,
			Author:   "-",
			Message:  watchEvent.Message,
			Date:     watchEvent.EventDate.Format("2006-01-02 15:04:05"),
			Name:     watchEvent.Name,
			Url:      watchEvent.Url,
			Platform: watchEvent.Platform,
		})
		if err != nil {
			log.Printf("Error sending message: %v %v", watchEvent, err)
		}
	}

	return nil
}
