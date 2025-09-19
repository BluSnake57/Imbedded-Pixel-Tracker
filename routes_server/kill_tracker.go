package routes_server

import (
	Database "IPT/Imbedded-Pixel-Tracker/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Kill_Tracker(gc *gin.Context) {
	var tracker Database.Tracker

	// Parses JSON received from client
	err := gc.ShouldBindJSON(&tracker)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = Database.Remove_Tracker(tracker.ClientID)
	if err != nil {
		log.Println("Internal Server Error: ", err)
		gc.JSON(http.StatusInternalServerError, err.Error())
	}

	gc.JSON(http.StatusOK, tracker)
}
