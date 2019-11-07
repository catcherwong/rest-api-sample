package biz

import (
	"github.com/catcherwong/rest-api-sample/data"
	"github.com/catcherwong/rest-api-sample/dto"
	models2 "github.com/catcherwong/rest-api-sample/models"
)

type UserBiz struct {
	userRepo *data.UserRepo
}

func NewUserBiz() *UserBiz {
	return &UserBiz{&data.UserRepo{}}
}

func (b UserBiz) CreateUser(u models2.User) error {

	err := b.userRepo.CreateUser(u)

	return err
}

func (b UserBiz) GetUserById(id int64) models2.User {

	u := b.userRepo.GetUserById(id)

	return u
}

func (b UserBiz) GetUserList(d *dto.GetUserListDto) []models2.User {

	if d.PageIndex <= 0 {
		d.PageIndex = 1
	}

	if d.PageSize <= 0 {
		d.PageSize = 10
	}

	l := b.userRepo.GetUserListByPage(d)

	return l
}
