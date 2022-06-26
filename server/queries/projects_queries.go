package queries

import (
	"github.com/zcubbs/pulse/server/models"
	"github.com/zcubbs/pulse/server/utils"
)

func GetAllProjects() *[]models.Project {
	return utils.Projects
}
