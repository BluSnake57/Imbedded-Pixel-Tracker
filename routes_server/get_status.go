package routes_server

import (
	Database "IPT/Imbedded-Pixel-Tracker/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_Status(gc *gin.Context) {
	var tracker Database.Tracker

	// Parses JSON received from client
	err := gc.ShouldBindJSON(&tracker)
	if err != nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tracker, err = Database.Tracker_Status(tracker.ClientID)
	if err != nil {
		log.Println("Tracker Record could not be found or recovered:", err)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gc.JSON(http.StatusOK, tracker)
}
