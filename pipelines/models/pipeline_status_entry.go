package models

import (
	"github.com/uptrace/bun"
	"time"
)

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
