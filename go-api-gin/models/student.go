package models

type Student struct {
	Id    string  `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Major string  `json:"major" binding:"required"`
	GPA   float64 `json:"gpa" binding:"gte=0,lte=4"`
}