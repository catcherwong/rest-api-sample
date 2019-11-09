package common

import "github.com/gin-gonic/gin"

type PageRequest struct {
	PageIndex int64 `form:"pageIndex"`
	PageSize  int64 `form:"pageSize"`
}

func GetPager(c *gin.Context) (PageRequest, error) {

	var req PageRequest

	err := c.ShouldBindQuery(&req)

	return req, err
}
