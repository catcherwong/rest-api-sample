package models

type User struct {
	Id     int64
	Name   string
	Gender int
}

func GetUserList() []User {

	l := []User{
		{Id: 1, Name: "catcher", Gender: 1},
		{Id: 2, Name: "wong", Gender: 1},
	}

	return l
}

func GetUserById(id int64) User {

	u := User{Id: id, Name: "catcher", Gender: 1}
	return u
}
