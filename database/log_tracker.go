package Database

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Log_Tracker(tracker Tracker) error {
	tracker_record, err := get_tracker(tracker.TrackerID)
	if err != nil {
		log.Println("Tracker Record not found:", err)
		return err
	}

	tracker.ClientID = tracker_record.ClientID
	tracker.Init_Time = tracker_record.Init_Time
	tracker.Access_Time = time.Now()

	err = set_tracker(tracker)
	if err != nil {
		log.Println("Tracker could not be saved:", err)
		return err
	}

	return nil
}

func get_tracker(trackerID string) (Tracker, error) {

	entries, err := os.ReadDir(Tracker_Records)
	if err != nil {
		log.Fatalf("Error reading directory %s: %v", Tracker_Records, err)
	}

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".json") {
			tracker_record, err := compare_trackerID_to_tracker_record(trackerID, entry.Name())
			if err != nil {
				return Tracker{}, err
			}
			if tracker_record.ClientID != "" {
				return tracker_record, err
			}
		}
	}

	return Tracker{}, errors.New("Tracker Record could not be found")
}

// Compares provided trackerID to a tracker record file
func compare_trackerID_to_tracker_record(trackerID string, entry string) (Tracker, error) {
	var tracker_record Tracker

	file, err := os.OpenFile(filepath.Join(Tracker_Records, entry), os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Printf("Error reading directory %s: %v", Tracker_Records, err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tracker_record)
	if err != nil {
		log.Println("Error decoding JSON:", err)
	}

	if tracker_record.TrackerID == trackerID {
		return tracker_record, err
	}

	return Tracker{}, err
}

func set_tracker(tracker Tracker) error {
	file, err := os.OpenFile(filepath.Join(Tracker_Records, tracker.ClientID+".json"), os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Printf("Error saving record %s: %v", Tracker_Records, err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(tracker)

	return nil
}
