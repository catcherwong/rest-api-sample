package controllers

import (
	"github.com/catcherwong/rest-api-sample/biz"
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/catcherwong/rest-api-sample/dto"
	models2 "github.com/catcherwong/rest-api-sample/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	userBiz *biz.UserBiz
}

func NewUserController() *UserController {
	return &UserController{biz.NewUserBiz()}
}

// @Summary Add User
// @Tags User
// @Produce  json
// @Param req body dto.UserReq true "Name"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/users [post]
func (uc UserController) AddUser(c *gin.Context) {

	var req dto.UserReq
	err := c.ShouldBindJSON(&req)

	if err != nil {
		log.Println("bind json error", err)
		c.JSON(http.StatusOK, common.GetErrorResponse("param error"))
		return
	}

	err = uc.userBiz.CreateUser(models2.User{Name: req.Name, Gender: req.Gender})

	if err != nil {
		log.Println("db error", err)
		c.JSON(http.StatusOK, common.GetErrorResponse("db error"))
		return
	}

	c.JSON(http.StatusOK, common.GetOKResponse(true))
}

// @Summary Get User By Id
// @Tags User
// @Produce  json
// @Param id path int true "Id"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/users/{id} [get]
func (uc UserController) GetUserById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || id <= 0 {
		c.JSON(http.StatusOK, common.GetErrorResponse("id error"))
		return
	}

	u := uc.userBiz.GetUserById(id)

	c.JSON(http.StatusOK, common.GetOKResponse(u))
}

// @Summary Get User List
// @Tags User
// @Produce  json
// @Param id query int false "Id"
// @Param name query string false "Name"
// @Param gender query int false "Gender"
// @Param pageIndex query int false "PageIndex"
// @Param pageSize query int false "PageSize"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/users [get]
func (uc UserController) GetUserList(c *gin.Context) {
	var dto dto.GetUserListDto

	err := c.ShouldBindQuery(&dto)

	if err != nil {
		c.JSON(http.StatusOK, common.GetErrorResponse("bind error"))
		return
	}

	l := uc.userBiz.GetUserList(&dto)

	c.JSON(http.StatusOK, common.GetOKResponse(l))

}
