package redis

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"github.com/catcherwong/bgadmin-rest-api/common"
	"github.com/catcherwong/bgadmin-rest-api/db"
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

	type SetStringRequest struct {
		Key        string `json:"key"`
		Value      string `json:"val"`
		Expiration int64  `json:"exp"`
	}

	var r SetStringRequest

	err := c.ShouldBindJSON(&r)

	if err != nil {
		log.Println("bind param error")

		d = common.GetErrorResponse("can not bind param")
		c.JSON(http.StatusOK, d)
		return
	}

	t := time.Second * time.Duration(r.Expiration)

	err = db.RedisClient.Set(r.Key, r.Value, t).Err()

	if err != nil {
		log.Println("bind param error")

		d = common.GetErrorResponse("set error")
		c.JSON(http.StatusOK, d)
		return
	}

	c.JSON(http.StatusOK, d)
}

func deleteValue(c *gin.Context) {

	d := common.GetOKResponse("")

	key, ok := c.GetQuery("key")

	if !ok {
		d = common.GetErrorResponse("key can not be empty")
		c.JSON(http.StatusOK, d)
		return
	}

	err := db.RedisClient.Del(key).Err()

	if err != nil {

		log.Println("del redis key [", err.Error(), "] error")

		d = common.GetErrorResponse("can not delete the value")
		c.JSON(http.StatusOK, d)
		return
	}

	c.JSON(http.StatusOK, d)
}

func getString(c *gin.Context) {

	d := common.GetOKResponse("")

	key, ok := c.GetQuery("key")

	if !ok {
		d = common.GetErrorResponse("key can not be empty")
		c.JSON(http.StatusOK, d)
		return
	}

	v, err := db.RedisClient.Get(key).Result()

	if err != nil {

		log.Println("get redis key [", err.Error(), "] error")

		d = common.GetErrorResponse("can not get the value")
		c.JSON(http.StatusOK, d)
		return
	}

	d.SetData(v)

	c.JSON(http.StatusOK, d)
}
