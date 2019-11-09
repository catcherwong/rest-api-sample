package controllers

import (
	"github.com/catcherwong/rest-api-sample/biz"
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/catcherwong/rest-api-sample/dto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RedisController struct {
	redisBiz *biz.RedisBiz
}

func NewRedisController() *RedisController {
	return &RedisController{biz.NewRedisBiz()}
}

// @Summary Set string value to redis
// @Tags Redis
// @Produce  json
// @Param req body dto.SetStringRequest true "Name"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/redis/string [post]
func (rc RedisController) SetString(c *gin.Context) {

	d := common.GetOKResponse("")

	var r dto.SetStringRequest

	err := c.ShouldBindJSON(&r)

	if err != nil {
		log.Println("bind param error")

		d = common.GetErrorResponse("can not bind param")
		c.JSON(http.StatusOK, d)
		return
	}

	_, err = rc.redisBiz.SetString(r.Key, r.Value, r.Expiration)

	if err != nil {
		log.Println("bind param error")

		d = common.GetErrorResponse("set error")
		c.JSON(http.StatusOK, d)
		return
	}

	c.JSON(http.StatusOK, d)
}

// @Summary Delete a key from redis
// @Tags Redis
// @Produce  json
// @Param  key query string true "Key"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/redis [delete]
func (rc RedisController) DeleteValue(c *gin.Context) {

	d := common.GetOKResponse("")

	key, ok := c.GetQuery("key")

	if !ok {
		d = common.GetErrorResponse("key can not be empty")
		c.JSON(http.StatusOK, d)
		return
	}

	_, err := rc.redisBiz.DeleteValue(key)

	if err != nil {

		log.Println("del redis key [", err.Error(), "] error")

		d = common.GetErrorResponse("can not delete the value")
		c.JSON(http.StatusOK, d)
		return
	}

	c.JSON(http.StatusOK, d)
}

// @Summary Get a string value from redis
// @Tags Redis
// @Produce  json
// @Param  key query string true "Key"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/redis/string [get]
func (rc RedisController) GetString(c *gin.Context) {

	d := common.GetOKResponse("")

	key, ok := c.GetQuery("key")

	if !ok {
		d = common.GetErrorResponse("key can not be empty")
		c.JSON(http.StatusOK, d)
		return
	}

	v, err := rc.redisBiz.GetString(key)

	if err != nil {

		log.Println("get redis key [", err.Error(), "] error")

		d = common.GetErrorResponse("can not get the value")
		c.JSON(http.StatusOK, d)
		return
	}

	d.SetData(v)

	c.JSON(http.StatusOK, d)
}
