package models

type App struct {
	Name  string `json:"Name"`
	AppID string `json:"AppID"`
}

type Config struct {
	Apps   []App            `json:"apps"`
	Groups map[string][]App `json:"groups"`
}
