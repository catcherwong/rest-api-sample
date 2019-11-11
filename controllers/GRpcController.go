package controllers

import (
	"context"
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/catcherwong/rest-api-sample/svcs"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type GRpcController struct {
}

func NewGRpcController() *GRpcController {
	return &GRpcController{}
}

// @Summary Test GRPC
// @Tags GRPC
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/grpc [get]
func (rc GRpcController) GetString(c *gin.Context) {

	conn, err := grpc.Dial(":19999", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	client := svcs.NewGreeterClient(conn)

	name, _ := c.GetQuery("name")
	log.Println(name)
	result, err := client.SayHello(context.Background(), &svcs.HelloRequest{Name: name})

	if err != nil {
		log.Println("request error")
	}
	log.Println(result.Message)

	c.JSON(http.StatusOK, common.GetOKResponse(result.Message))
}
