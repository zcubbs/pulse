package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/zcubbs/pulse/models"
	"github.com/zcubbs/pulse/utils"
	"log"
	"strconv"
)

const objectBuild string = "build"

func HandleGitlabPipelineEvent(c *fiber.Ctx) error {
	event := c.Get("X-Gitlab-Event")
	if len(event) == 0 {
		log.Println(gitlab.ErrMissingGitLabEventHeader)
		return c.SendStatus(500)
	}
	gitLabEvent := gitlab.Event(event)
	payload := c.Body()
	parsedPayload, err := eventParsing(gitLabEvent, payload, gitlab.PipelineEvents, gitlab.JobEvents)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	log.Println(parsedPayload)
	err = processGitlabPayload(parsedPayload)
	if err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}
	return c.SendStatus(200)
}

func processGitlabPayload(payload interface{}) error {
	switch payload.(type) {

	case gitlab.PipelineEventPayload:
		pipeline := payload.(gitlab.PipelineEventPayload)

		entry := &models.PipelineStatusEntry{
			Origin:      "Gitlab",
			OriginUrl:   pipeline.Project.WebURL,
			Status:      pipeline.ObjectAttributes.Status,
			Message:     fmt.Sprintf("By user %s", pipeline.User.Name),
			ProjectId:   strconv.FormatInt(pipeline.Project.ID, 16),
			ProjectName: pipeline.Project.Name,
			Group:       pipeline.Project.Namespace,
		}
		fmt.Printf("%+v", pipeline)
		fmt.Printf("\n%v\n", entry)
		mEntry, err := json.Marshal(entry)
		if err != nil {
			log.Println(err)
			return err
		}
		// Create a message to publish.
		message := amqp.Publishing{
			ContentType: "application/json",
			Body:        mEntry,
		}

		log.Println("received: ", message.Body)

		// Attempt to publish a message to the queue.
		publishEvent(message)

	case gitlab.JobEventPayload:
		job := payload.(gitlab.JobEventPayload)
		fmt.Printf("%+v", job)
	}

	return nil
}

func publishEvent(message amqp.Publishing) {
	if err := utils.GetChannelRabbitMQ().Publish(
		"",                                     // exchange
		"zrocket_pipeline_status_events.input", // queue name
		false,                                  // mandatory
		false,                                  // immediate
		message,                                // message to publish
	); err != nil {
		log.Println(err)
	}
}

func eventParsing(gitLabEvent gitlab.Event, payload []byte, events ...gitlab.Event) (interface{}, error) {

	var found bool
	for _, evt := range events {
		if evt == gitLabEvent {
			found = true
			break
		}
	}
	// event not defined to be parsed
	if !found {
		return nil, gitlab.ErrEventNotFound
	}

	switch gitLabEvent {
	case gitlab.PipelineEvents:
		var pl gitlab.PipelineEventPayload
		err := json.Unmarshal(payload, &pl)
		return pl, err

	case gitlab.JobEvents:
		var pl gitlab.JobEventPayload
		err := json.Unmarshal(payload, &pl)
		if err != nil {
			return nil, err
		}
		if pl.ObjectKind == objectBuild {
			return eventParsing(gitlab.BuildEvents, payload, events...)
		}
		return pl, nil
	default:
		return nil, fmt.Errorf("unknown event %s", gitLabEvent)
	}
}
