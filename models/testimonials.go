package models

type Testimonial struct {
    Name  string `json:"name"`
    Job string `json:"job"`
    Comment string `json:"comment"`
    Slug string `json:"slug"`
}

type Testimonials struct {
	Testimonials			[]Testimonial `json:"testimonials"`
}