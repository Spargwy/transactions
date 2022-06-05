package model

type User struct {
	ID   int     `json:"ID"`
	Name *string `json:"Name"`
	//Возможно, стоит делать uint'ом
	Balance int `json:"Balance"`
}
