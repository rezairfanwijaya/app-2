package model

type company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Companies = []company{
	{ID: "1", Name: "TechCorp"},
	{ID: "2", Name: "HealthSolutions"},
	{ID: "3", Name: "FinAdvisors"},
	{ID: "4", Name: "EduPro"},
	{ID: "5", Name: "MediCare Group"},
	{ID: "6", Name: "EcoGreen Industries"},
	{ID: "7", Name: "FoodieDelight"},
	{ID: "8", Name: "TravelGo"},
	{ID: "9", Name: "AutoWorld"},
	{ID: "10", Name: "BuildRight Construction"},
}
