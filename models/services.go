package models

type Service struct {
	Title			string `json:"title"`
	Description 	string `json:"description"`
}

type Services struct {
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Services			[]Service `json:"services"`
}