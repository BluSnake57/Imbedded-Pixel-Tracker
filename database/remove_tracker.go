package Database

import (
	"log"
	"os"
	"path/filepath"
)

func Remove_Tracker(clientID string) error {
	err := os.Remove(filepath.Join(Tracker_Records, clientID+".json"))

	if err != nil {
		log.Println("Tracker could not be removed")
	}

	return err
}
