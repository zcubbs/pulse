package utils

import (
	"github.com/zcubbs/pulse/server/models"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var Projects *[]models.Project

func LoadYamlConfig() {
	yamlFile, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	yamlProjects := &models.Projects{}

	err = yaml.Unmarshal(yamlFile, yamlProjects)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(yamlProjects.Projects)
	Projects = yamlProjects.Projects
}

func getConfigPath() string {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	return configPath
}
