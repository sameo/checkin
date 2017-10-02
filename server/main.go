package main

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func checkinTimeStamp(w http.ResponseWriter, r *http.Request) {
	var containerID, event string
	
	r.ParseForm()
	ts := time.Now()
	tsString := ts.Format(time.RFC3339Nano)
	for k, v := range r.Form {
		switch k {
		case "containerID":
			containerID = v[0]
		case "event":
			event = v[0]
		}
	}

	if event != "" && containerID != "" {
		log.WithFields(log.Fields{
			"Event": event,
			"Container ID":   containerID,
		}).Info(tsString)
	}
}

func main() {
	log.SetLevel(log.InfoLevel)
	log.Info("Starting checkin server...")
	http.HandleFunc("/checkin", checkinTimeStamp) // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
