package data

import "my-fiber-app/models"

var Home = models.Home{
	Title: "Hello, I’mBrooklyn Gilbert",
	Description: "I'm a Freelance UI/UX Designer and Developer based in London, England. I strives to build immersive and beautiful web applications through carefully crafted code and user-centric design.",
	Experience: []models.Exp{
		{Name: "Experience", Value: "15 Y."},
		{Name: "Projects Completed", Value: "250+"},
		{Name: "Happy Clients", Value: "58"},
	},
	Img: "/data/img",
}

var About = models.About{
	Img: "/data/img/hero-bg.jpg",
	Title: "I am Professional User Experience Designer",
	Description: "I design and develop services for customers specializing creating stylish, modern websites, web services and online stores. My passion is to design digital user experiences.I design and develop services for customers specializing creating stylish, modern websites, web services.",
	Links: []models.Link{
		{Name: "CV", Value: "#"},
		{Name: "Facebook", Value: "facebook.com"},
		{Name: "Instagram", Value: "instagram.com"},
		{Name: "LinkedIn", Value: "linkedin.com"},
	},
}

