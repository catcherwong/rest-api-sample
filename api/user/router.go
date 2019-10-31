package user

import (
	"github.com/catcherwong/bgadmin-rest-api/common"
	"github.com/catcherwong/bgadmin-rest-api/db/models"
	"github.com/gin-gonic/gin"
	"net/http"

	"strconv"
)

func InitRouters(r *gin.RouterGroup) {

	rg := r.Group("/users")
	{
		rg.GET("/", getUserList)
		rg.GET("/:id", getUserById)
	}
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
