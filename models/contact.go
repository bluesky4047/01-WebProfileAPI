package models

type Contact struct {
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Details			[]Detail `json:"details"`
}