package models

import (
	"fmt"
	"github.com/catcherwong/rest-api-sample/db"
	"github.com/catcherwong/rest-api-sample/dto"
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

func GetUserListByPage(d *dto.GetUserListDto) []User {

	if d.PageIndex <= 0 {
		d.PageIndex = 1
	}

	if d.PageSize <= 0 {
		d.PageSize = 10
	}

	limit := (d.PageIndex - 1) * d.PageSize

	var l []User

	var args []interface{}

	sql := "select id, name, gender from userinfo where 1 = 1 "

	if d.Id > 0 {
		sql += " and id = ?"
		args = append(args, d.Id)
	}

	if d.Name != "" {
		sql += " and name = ?"
		args = append(args, d.Name)
	}

	if d.Gender > 0 {
		sql += " and gender = ?"
		args = append(args, d.Gender)
	}

	sql += fmt.Sprintf(" limit %d offset %d ", d.PageSize, limit)

	log.Print(sql)

	rows, err := db.DB.Query(sql, args...)

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
