package repository

//User type
type User struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

// Data In built DB
var Data = []User{
	{Name: "Onur", Balance: -99},
	{Name: "Atilla", Balance: 1000},
	{Name: "Nazim", Balance: 999},
	{Name: "Abdulsamet", Balance: 999},
	{Name: "Deniz", Balance: 999},
}
