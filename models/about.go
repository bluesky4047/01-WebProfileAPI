package models

type Link struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

type About struct {
	Img				string `json:"img"`
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Links			[]Link `json:"links"`
}