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

var Process = models.Process{
	Title: "Work Process",
	Description: "Driven by design and powered by code, I create digital interfaces that feel intuitive and perform seamlessly. Every layout, animation, and component is crafted with intention — merging usability with visual clarity, I blend clean design with efficient code to build engaging, user-friendly web experiences that stand out.",
	Details: []models.Detail{
		{Name: "1. Research", Description: "Design meets function in every pixel, blending clarity with intuitive motion."},
		{Name: "2. Analyze", Description: "Crafting clean, thoughtful interfaces where form flows seamlessly into function and clarity."},
		{Name: "3. Design", Description: "I design seamless digital experiences with precision, purpose, and a touch of elegance."},
		{Name: "4. Launch", Description: "I craft digital products where thoughtful design meets performance-driven, responsive development."},
	},
}

