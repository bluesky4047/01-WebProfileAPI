package models

type Exp struct {
    Name  string `json:"name"`
    Value string `json:"value"`
}

type Home struct {
	Title	string `json:"title"`
	Description	string `json:"description"`
	Experience	[]Exp `json:"experience"`
	Img		string `json:"img"`
}