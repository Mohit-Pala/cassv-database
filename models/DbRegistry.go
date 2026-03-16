package models

type DbRegistry struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FolderPath string `json:"folder_path"`
}
