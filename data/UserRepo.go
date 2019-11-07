package data

import (
	"github.com/catcherwong/rest-api-sample/db"
	"github.com/catcherwong/rest-api-sample/dto"
	"github.com/catcherwong/rest-api-sample/models"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r UserRepo) GetUserListByPage(d *dto.GetUserListDto) []models.User {

	var l []models.User

	limit := (d.PageIndex - 1) * d.PageSize

	conn := db.SQLiteDb.Table("userinfo")

	if d.Id > 0 {
		conn = conn.Where(" id = ? ", d.Id)
	}

	if d.Name != "" {
		conn = conn.Where(" name = ? ", d.Name)
	}

	if d.Gender > 0 {
		conn = conn.Where(" gender = ?", d.Gender)
	}

	conn.Limit(d.PageSize).Offset(limit).Find(&l)

	return l
}

func (r UserRepo) GetUserById(id int64) models.User {

	var u models.User

	db.SQLiteDb.Table("userinfo").Where("id = ? ", id).First(&u)

	return u
}

func (r UserRepo) CreateUser(u models.User) error {

	err := db.SQLiteDb.Raw("insert into userinfo (name, gender) values (?,?)", u.Name, u.Gender).Error

	return err
}
