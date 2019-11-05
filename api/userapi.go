package api

import (
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/catcherwong/rest-api-sample/db/models"
	"github.com/catcherwong/rest-api-sample/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type userApi struct {
}

func NewUserApi() *userApi {
	return &userApi{}
}

func (ua userApi) InitRouter(r *gin.Engine) {

	rg := r.Group("/api/users")
	{
		rg.GET("/", ua.GetUserList)
		rg.GET("/:id", ua.GetUserById)
		rg.POST("/", ua.AddUser)
	}
}

func (ua userApi) AddUser(c *gin.Context) {

	type UserReq struct {
		Name   string `json:"name"`
		Gender int    `json:"gender"`
	}
	var req UserReq
	err := c.ShouldBindJSON(&req)

	if err != nil {
		log.Println("bind json error", err)
		c.JSON(http.StatusOK, common.GetErrorResponse("param error"))
		return
	}

	err = models.CreateUser(models.User{Name: req.Name, Gender: req.Gender})

	if err != nil {
		log.Println("db error", err)
		c.JSON(http.StatusOK, common.GetErrorResponse("db error"))
		return
	}

	c.JSON(http.StatusOK, common.GetOKResponse("ok"))
}

func (ua userApi) GetUserById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || id <= 0 {
		c.JSON(http.StatusOK, common.GetErrorResponse("id error"))
		return
	}

	u := models.GetUserById(id)

	c.JSON(http.StatusOK, common.GetOKResponse(u))
}

func (ua userApi) GetUserList(c *gin.Context) {
	var dto dto.GetUserListDto

	err := c.ShouldBindQuery(&dto)

	if err != nil {
		c.JSON(http.StatusOK, common.GetErrorResponse("bind error"))
		return
	}

	l := models.GetUserListByPage(&dto)

	c.JSON(http.StatusOK, common.GetOKResponse(l))

}
