package controllers

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/zcubbs/pulse/server/models"
	"github.com/zcubbs/pulse/server/queries"
	"github.com/zcubbs/pulse/server/utils"
	"log"
	"sort"
)

func HandleGetWatchReports(c *fiber.Ctx) error {
	watchEvents := queries.GetProjectLatestPipelineStatusEntries()
	sort.Sort(models.ByEventDate(watchEvents))
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
