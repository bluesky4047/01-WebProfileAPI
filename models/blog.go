package models

type Blog struct {
	Img				string `json:"img"`
	Title			string `json:"title"`
	Date			string `json:"date"`
	TotalComments		string `json:"totalComments"`
}

type Blogs struct {
	Blogs			[]Blog `json:"blogs"`
}