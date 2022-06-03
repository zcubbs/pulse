package queries

import (
	"github.com/zcubbs/pulse/models"
	"github.com/zcubbs/pulse/utils"
)

func GetAllProjects() *[]models.Project {
	return utils.Projects
}
