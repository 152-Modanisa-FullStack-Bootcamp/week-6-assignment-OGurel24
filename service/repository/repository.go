package repository

type User struct {
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}



var Data = []User{
	{Name: "onur", Balance: 999},
	{Name: "ugur", Balance: 10},
}
