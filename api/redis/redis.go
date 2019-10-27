package redis

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/catcherwong/bgadmin-rest-api/common"
)

// InitRouters inits redis route
func InitRouters(r *gin.RouterGroup) {

	rg := r.Group("/redis")
	{
		rg.GET("/string", getString)
		rg.POST("/string", setString)
		rg.DELETE("/", deleteValue)
	}
}

func setString(c *gin.Context) {

	d := common.GetOKResponse("")

	c.JSON(http.StatusOK, d)
}

func deleteValue(c *gin.Context) {

	d := common.GetOKResponse("")

	c.JSON(http.StatusOK, d)
}

func getString(c *gin.Context) {

	d := common.GetOKResponse("")

	c.JSON(http.StatusOK, d)
}
