package Database

import "time"

const (
	Tracking_Pixel  = "tracking-pixel.png"
	Tracker_Records = "tracker_records/"
)

type Tracker struct {
	ClientID      string    `json:"clientID"`
	TrackerID     string    `json:"trackerID"`
	Init_Time     time.Time `json:"init-time"`
	IP_Address    string    `json:"RemoteIP"`
	CF_IP_Address string    `json:"CF-Connecting-IP"`
	CF_IP_Path    string    `json:"X-Forwarded-For"`
	CF_Country    string    `json:"CF-IPCountry"`
	Accessed      bool      `json:"accessed"`
	Access_Time   time.Time `json:"access-time"`
}
