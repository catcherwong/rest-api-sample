package controllers

import (
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type MetricsController struct {
}

func init() {
	prometheus.MustRegister(bizTotal)
}

func NewMetricsController() *MetricsController {
	return &MetricsController{}
}

// @Summary Record some metrics
// @Tags Metrics
// @Produce  json
// @Success 200 {object} common.RestResponse "{"code":0,"data":{},"msg":"ok"}"
// @Router /api/v1/metrics [get]
func (mc MetricsController) Record(c *gin.Context) {

	bizTotal.With(prometheus.Labels{"biztype": "mybiztype", "bizname": "mybizname"}).Inc()

	c.JSON(http.StatusOK, common.GetOKResponse("ok"))
}

var bizTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "rest_api_sample_biz",
		Help: "Number of hello requests in total",
	},
	[]string{"biztype", "bizname"},
)
