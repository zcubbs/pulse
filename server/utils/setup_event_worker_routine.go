package utils

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/zcubbs/pulse/server/models"
	"log"
	"time"
	_ "time/tzdata"
)

func LaunchEventWorker() {
	// Subscribing to QueueService1 for getting messages.
	messages, err := channelRabbitMQ.Consume(
		"zrocket_pipeline_status_events.input", // queue name
		"",                                     // consumer
		true,                                   // auto-ack
		false,                                  // exclusive
		false,                                  // no local
		false,                                  // no wait
		nil,                                    // arguments
	)
	if err != nil {
		log.Println(err)
	}

	go func() {
		t := time.Tick(5 * time.Second)
		for {
			select {
			case <-t:
				processMessage(messages)
			case <-ctx.Done():
				return
			}
		}
	}()

	log.Println("Successfully started event worker routine")
}

func processMessage(messages <-chan amqp.Delivery) {
	for message := range messages {
		entry := &models.PipelineStatusEntry{}
		err2 := json.Unmarshal(message.Body, entry)
		if err2 != nil {
			log.Println(err2)
		}

		insertWatchEntry(entry)
	}
}

func insertWatchEntry(entry *models.PipelineStatusEntry) {
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		log.Println(err)
	}
	now := time.Now().In(loc)
	entry.EventDate = now

	_, err = GetPgDatabase().NewInsert().
		Model(entry).
		On("CONFLICT (project_id) DO UPDATE").
		Set("status = EXCLUDED.status").
		Set("event_date = EXCLUDED.event_date").
		Exec(ctx)

	if err != nil {
		log.Println(err)
	}

	notifyEdgeService(&models.WSMessage{
		MessageType: "watch_refresh",
		Message:     fmt.Sprintf("Project %s status changed to %s", entry.ProjectName, entry.Status),
	})
}

func notifyEdgeService(message *models.WSMessage) {
	WriteMessage(message)
}

// External call => Not needed for now
//func notifyEdgeService(message WSMessage) {
//	messageJSON, err := json.Marshal(message)
//	if err != nil {
//		log.Println(err)
//	}
//
//	_, err = http.Post(
//		fmt.Sprintf("%s%s", getEdgeServiceUrl(), "/api/v1/watch/wsRefresh"),
//		"application/json",
//		bytes.NewBuffer(messageJSON),
//	)
//
//	if err != nil {
//		log.Println(err)
//	}
//}
