package Database

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func Generate_Tracker(clientID string) (string, error) {
	var tracker Tracker

	tracker.ClientID = clientID
	//Change trackerID to be based off of a hash of clientID
	//This would make it easier to verify both ClientID and TrackerID without needing to open the record
	tracker.TrackerID = GenerateRandomStringURLSafe(16) //Generates a random number with length of 16 bytes
	tracker.Init_Time = time.Now()
	tracker.IP_Address = ""
	tracker.CF_IP_Address = ""
	tracker.CF_IP_Path = ""
	tracker.CF_Country = ""
	tracker.Accessed = false
	tracker.Access_Time = time.Now()

	err := generate_tracker_record(tracker)
	if err != nil {
		log.Println("Tracker Record already exists:", err)
	}

	tracker_url := generate_tracker_url(tracker)

	return tracker_url, err
}

// generate_tracker_record validates and stores the new tracker
func generate_tracker_record(tracker Tracker) error {
	log.Println(os.ReadDir(Tracker_Records))

	//Checks for other records of same ClientID
	entries, err := os.ReadDir(Tracker_Records)
	if err != nil {
		log.Fatalf("Error reading directory %s: %v", Tracker_Records, err)
	}

	for _, entry := range entries {
		if entry.Name() == tracker.ClientID+".json" {
			return errors.New("ClientID:" + tracker.ClientID + " already exists")
		}
	}

	//Creates record and saves data
	file, err := os.OpenFile(filepath.Join(Tracker_Records, tracker.ClientID+".json"), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalln("Could not generate Tracker record:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(tracker)

	return err
}

// Generates the URL for the tracker
func generate_tracker_url(tracker Tracker) string {
	var new_url = os.Getenv("PROTOCOL") + "://" + os.Getenv("HOST") + "/trace_client" + "?" + os.Getenv("CLIENT_VAR") + "=" + tracker.TrackerID

	u, err := url.ParseRequestURI(new_url)
	if err != nil {
		log.Fatalln("Could not generate valid URL:", err)
	}

	return u.String()
}
