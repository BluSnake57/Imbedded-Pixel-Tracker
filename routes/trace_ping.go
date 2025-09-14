package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingTest struct {
	Ping string `json:"ping"`
}

func Trace_Ping(gc *gin.Context) {
	println(gc.RemoteIP())
	fmt.Println("request header", gc.Request.Header)
	body, _ := io.ReadAll(gc.Request.Body)
	println(string(body))
	pingTest := PingTest{
		Ping: "pong",
	}

	gc.JSON(http.StatusOK, pingTest)
}
