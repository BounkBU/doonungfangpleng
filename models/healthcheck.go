package model

type ServerInformation struct {
	Project      string   `json:"project"`
	Team         string   `json:"team"`
	Collaborator []string `json:"collaborator"`
}
