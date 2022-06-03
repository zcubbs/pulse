package queries

import (
	"context"
	"github.com/zcubbs/pulse/models"
	"github.com/zcubbs/pulse/utils"
	"log"
	"time"
)

var ctx = context.Background()

func GetProjectLatestPipelineStatusEntryForProject(projectId string) []models.PipelineStatusEntry {
	pipelineStatusEntries := make([]models.PipelineStatusEntry, 0)
	if err := utils.GetPgDatabase().NewSelect().Model(&pipelineStatusEntries).Where("project_id LIKE ?", projectId).Scan(ctx); err != nil {
		log.Println(err)
	}
	return pipelineStatusEntries
}

func GetProjectLatestPipelineStatusEntries() []models.PipelineStatusEntry {
	pipelineStatusEntries := make([]models.PipelineStatusEntry, 0)
	if err := utils.GetPgDatabase().NewSelect().Model(&pipelineStatusEntries).Scan(ctx); err != nil {
		log.Println(err)
	}
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		log.Println(err)
	}

	pipelineStatusEntriesWithTZ := make([]models.PipelineStatusEntry, 0)
	for _, v := range pipelineStatusEntries {
		v.EventDate = v.EventDate.In(loc)
		pipelineStatusEntriesWithTZ = append(pipelineStatusEntriesWithTZ, v)
	}
	return pipelineStatusEntriesWithTZ
}
