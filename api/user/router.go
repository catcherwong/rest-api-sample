package user

import (
	"github.com/catcherwong/bgadmin-rest-api/common"
	"github.com/catcherwong/bgadmin-rest-api/db/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"strconv"
)

func InitRouters(r *gin.RouterGroup) {

	rg := r.Group("/users")
	{
		rg.GET("/", getUserList)
		rg.GET("/:id", getUserById)
		rg.POST("/", addUser)
	}
}

func addUser(c *gin.Context) {

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

func getUserById(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || id <= 0 {
		c.JSON(http.StatusOK, common.GetErrorResponse("id error"))
		return
	}

	u := models.GetUserById(id)

	c.JSON(http.StatusOK, common.GetOKResponse(u))
}

func getUserList(c *gin.Context) {

	l := models.GetUserList()

	c.JSON(http.StatusOK, common.GetOKResponse(l))

}
