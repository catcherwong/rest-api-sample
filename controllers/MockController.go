package controllers

import (
	"github.com/catcherwong/rest-api-sample/biz"
	"github.com/catcherwong/rest-api-sample/common"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type MockController struct {
	redisBiz *biz.RedisBiz
}

func NewMockController() *MockController {
	return &MockController{}
}

// @Summary Get some result after 1 second
// @Tags Mock
// @Produce  plain
// @Success 200 {string} string "1"
// @Router /api/mock/get1 [get]
func (rc MockController) GetString1(c *gin.Context) {

	time.Sleep(time.Second)

	c.String(http.StatusOK, "1")
}

// @Summary Get some result after 2 second
// @Tags Mock
// @Produce  plain
// @Success 200 {string} string "2"
// @Router /api/mock/get2 [get]
func (rc MockController) GetString2(c *gin.Context) {

	time.Sleep(2 * time.Second)

	c.String(http.StatusOK, "2")
}

// @Summary Get some result after 2 second
// @Tags Mock
// @Produce  json
// @Success 200 {object} common.RestResponse "{"data":[]}"
// @Router /api/mock [get]
func (rc MockController) GetString(c *gin.Context) {

	urls := []string{"http://localhost:9999/api/mock/get1", "http://localhost:9999/api/mock/get2"}

	ch := make(chan string)

	for _, v := range urls {
		go doRequest(v, ch)
	}

	res := make([]string, 0)

	for range urls {
		d := <-ch
		log.Println(d)
		res = append(res, d)
	}

	close(ch)

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func doRequest(url string, ch chan string) {

	log.Println(url)

	res := ""

	resp, err := common.HttpClient.Get(url)

	if err != nil {
		log.Println("request error", err.Error())
		ch <- res
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println("read error", err.Error())
		ch <- res
		return
	}

	res = string(body)
	ch <- res
}
