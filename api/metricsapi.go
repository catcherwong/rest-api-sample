package api

import (
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

type metricsApi struct {
}

func NewMetricsApi() *metricsApi {
	return &metricsApi{}
}

func init() {
	prometheus.MustRegister(bizTotal)
}

func (ua metricsApi) InitRouter(r *gin.Engine) {

	rg := r.Group("/api/metrics")
	{
		rg.GET("/", ua.Record)
	}
}

func (ua metricsApi) Record(c *gin.Context) {

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
