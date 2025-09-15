package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type PingTest struct {
	Ping string `json:"ping"`
}

func Trace_Ping(gc *gin.Context) {

	println(gc.Query(os.Getenv("CLIENT_VAR")))
	fmt.Println("request header", gc.Request)
	println(gc.RemoteIP())
	println(gc.GetHeader("CF-Connecting-IP"))
	println(gc.GetHeader("X-Forwarded-For"))
	println(gc.GetHeader("CF-IPCountry"))

	pingTest := PingTest{
		Ping: "pong",
	}

	gc.JSON(http.StatusOK, pingTest)
}
