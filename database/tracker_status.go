package Database

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func Tracker_Status(ClientID string) (Tracker, error) {
	var tracker Tracker

	file, err := os.OpenFile(filepath.Join(Tracker_Records, ClientID+".json"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("Error reading directory %s: %v", Tracker_Records, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tracker)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}

	return tracker, err
}
