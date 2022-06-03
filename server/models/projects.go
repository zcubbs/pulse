package models

type Projects struct {
	Projects *[]Project `json:"projects" yaml:"projects"`
}

type Project struct {
	Name     string           `json:"name" yaml:"name"`
	Emoji    string           `json:"emoji" yaml:"emoji"`
	Id       int32            `json:"id" yaml:"id"`
	GitUrl   string           `json:"git_url" yaml:"git_url"`
	Branches *[]ProjectBranch `json:"branches" yaml:"branches"`
}

type ProjectBranch struct {
	Name string `json:"name" yaml:"name"`
}
