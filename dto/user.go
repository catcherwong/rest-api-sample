package dto

import (
	"github.com/catcherwong/rest-api-sample/common"
)

type GetUserListDto struct {
	common.PageRequest

	Id     int64  `form:"id"`
	Name   string `form:"name"`
	Gender int    `form:"gender"`
}

type UserReq struct {
	Name   string `json:"name"`
	Gender int    `json:"gender"`
}
