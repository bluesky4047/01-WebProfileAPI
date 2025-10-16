package models

type Detail struct {
	Name			string `json:"name"`
	Description 	string `json:"description"`
}

type Process struct {
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Details			[]Detail `json:"details"`
}