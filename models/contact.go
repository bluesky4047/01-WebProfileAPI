package models

type Contact struct {
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Details			[]Detail `json:"details"`
}

type ContactRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Location string `json:"location"`
	Budget   string `json:"budget"`
	Subject  string `json:"subject"`
	Message  string `json:"message"`
}