package controllers

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
	"github.com/zcubbs/pulse/pipelines/models"
	"github.com/zcubbs/pulse/pipelines/proto/pipelines"
	"github.com/zcubbs/pulse/pipelines/queries"
	"github.com/zcubbs/pulse/server/utils"
	"io"
	"log"
	"time"
)

func HandleGetWatchReports(c *fiber.Ctx) error {
	req := &pipelines.GetStatusRequest{Group: "main"}
	stream, err := utils.PipelinesGrpcClient.GetStatus(context.Background(), req)
	if err != nil {
		log.Println("Failed to get gRPC response:", err)
	}

	var watchEvents []PipelineStatusEntry

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("HandleGetWatchReports: stream completed", err)
			break
		}
		if err != nil {
			log.Println("Error while reading the stream:", err)
		}

		date, _ := time.Parse("2006-01-02 15:04:05", msg.Date)

		watchEvents = append(watchEvents, PipelineStatusEntry{
			BaseModel: bun.BaseModel{},
			ID:        0,
			Branch:    msg.Branch,
			Platform:  msg.Platform,
			Url:       msg.Url,
			Status:    msg.Status,
			Message:   msg.Message,
			Commit:    msg.Commit,
			Id:        "",
			Name:      msg.Name,
			Group:     msg.Name,
			EventDate: date,
		})
	}
	return c.JSON(watchEvents)
}

func HandleGetLatestWatchReport(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	return c.JSON(queries.GetProjectLatestPipelineStatusEntryForProject(projectId))
}

func HandleNotifyRefresh(c *fiber.Ctx) error {
	message := &models.WSMessage{}
	err := json.Unmarshal(c.Body(), message)
	if err != nil {
		log.Println(err)
	}
	utils.WriteMessage(message)
	return c.SendStatus(200)
}

type PipelineStatusEntry struct {
	bun.BaseModel `bun:"pipeline_status_entries,alias:pse"`
	//ID            uuid.UUID `bun:",default:gen_random_uuid(),type:uuid,pk"`
	ID        int64
	Branch    string    `json:"branch" bun:"branch"`
	Platform  string    `json:"platform" bun:"platform"`
	Url       string    `json:"url" bun:"url"`
	Status    string    `json:"status" bun:"status"`
	Message   string    `json:"message" bun:"message"`
	Commit    string    `json:"commit" bun:"commit"`
	Id        string    `json:"id" bun:"project_id,unique"`
	Name      string    `json:"name" bun:"name"`
	Group     string    `json:"group" bun:"group"`
	EventDate time.Time `json:"event_date" bun:",nullzero,notnull,default:current_timestamp"`
}

type ByEventDate []PipelineStatusEntry

func (a ByEventDate) Len() int { return len(a) }
func (a ByEventDate) Less(i, j int) bool {
	return a[i].EventDate.Before(a[j].EventDate)
}
func (a ByEventDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
