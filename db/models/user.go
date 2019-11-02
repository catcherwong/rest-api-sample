package models

import (
	"github.com/catcherwong/bgadmin-rest-api/db"
	"log"
)

type User struct {
	Id     int64
	Name   string
	Gender int
}

func GetUserList() []User {

	var l []User

	sql := "select id, name, gender from userinfo"

	rows, err := db.DB.Query(sql)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Gender)

		l = append(l, u)
	}

	defer rows.Close()

	return l
}

func GetUserById(id int64) User {

	sql := "select id, name, gender from userinfo where id = ? limit 1"

	row := db.DB.QueryRow(sql, id)

	var u User
	row.Scan(&u.Id, &u.Name, &u.Gender)

	return u
}

func CreateUser(u User) error {

	sql := "insert into userinfo (name, gender) values (?,?)"

	s, err := db.DB.Prepare(sql)

	if err != nil {
		return err
	}

	res, err := s.Exec(u.Name, u.Gender)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()

	return err
}
