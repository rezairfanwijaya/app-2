package model

type employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Employees = []employee{
	{ID: "1", Name: "John Doe"},
	{ID: "2", Name: "Jane Smith"},
	{ID: "3", Name: "Michael Johnson"},
	{ID: "4", Name: "Emily Davis"},
	{ID: "5", Name: "William Brown"},
	{ID: "6", Name: "Olivia Garcia"},
	{ID: "7", Name: "James Miller"},
	{ID: "8", Name: "Isabella Wilson"},
	{ID: "9", Name: "Benjamin Martinez"},
	{ID: "10", Name: "Sophia Anderson"},
	{ID: "11", Name: "Alexander Thomas"},
	{ID: "12", Name: "Charlotte Lee"},
	{ID: "13", Name: "Daniel Harris"},
	{ID: "14", Name: "Mia Clark"},
	{ID: "15", Name: "Henry Lewis"},
	{ID: "16", Name: "Amelia Robinson"},
	{ID: "17", Name: "Samuel Walker"},
	{ID: "18", Name: "Harper Young"},
	{ID: "19", Name: "David King"},
	{ID: "20", Name: "Evelyn Scott"},
}
