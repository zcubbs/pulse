package queries

import (
	"github.com/zcubbs/pulse/pipelines/models"
	"github.com/zcubbs/pulse/pipelines/utils"
)

func GetAllProjects() *[]models.Project {
	return utils.Projects
}
