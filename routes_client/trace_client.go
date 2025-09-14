package routes_client

import (
	Database "IPT/Imbedded-Pixel-Tracker/database"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Trace_Client(gc *gin.Context) {
	var tracker Database.Tracker

	// println(gc.Query(os.Getenv("CLIENT_VAR")))
	// fmt.Println("request header", gc.Request)
	// println(gc.RemoteIP())
	// println(gc.GetHeader("CF-Connecting-IP"))
	// println(gc.GetHeader("X-Forwarded-For"))
	// println(gc.GetHeader("CF-IPCountry"))

	tracker.TrackerID = gc.Query(os.Getenv("CLIENT_VAR"))
	tracker.IP_Address = gc.RemoteIP()
	tracker.CF_IP_Address = gc.GetHeader("CF-Connecting-IP")
	tracker.CF_IP_Path = gc.GetHeader("X-Forwarded-For")
	tracker.CF_Country = gc.GetHeader("CF-IPCountry")
	tracker.Access_Time = time.Now()
	tracker.Accessed = true

	err := Database.Log_Tracker(tracker)
	if err != nil {
		log.Println("Error logging tracker info")
	}

	gc.File(Database.Tracking_Pixel)
}
