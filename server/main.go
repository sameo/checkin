package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
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
			"Event":        event,
			"Container ID": containerID,
		}).Info(tsString)
	}
}

func watchCheckinFiles(watcher *fsnotify.Watcher, rootDir string) {
	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Create == fsnotify.Create {
				ts := time.Now()
				tsString := ts.Format(time.RFC3339Nano)
				log.WithFields(log.Fields{
					"Event":        "Starting",
					"Container ID": event.Name,
				}).Info(tsString)
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				ts := time.Now()
				tsString := ts.Format(time.RFC3339Nano)
				log.WithFields(log.Fields{
					"Event":        "Running",
					"Container ID": event.Name,
				}).Info(tsString)
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
		}
	}
}

func main() {
	var httpServer bool
	var checkinDir string
	log.SetLevel(log.InfoLevel)

	flag.BoolVar(&httpServer, "http", false, "HTTP checkin")
	flag.StringVar(&checkinDir, "checkin-root", "/var/run/checkin-clear-containers",
		"Checkin root directory")

	flag.Parse()

	log.Info("Starting checkin server...")
	if httpServer {
		http.HandleFunc("/checkin", checkinTimeStamp) // set router
		if err := http.ListenAndServe(":9090", nil); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		if err := os.MkdirAll(checkinDir, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal(err)
		}
		defer watcher.Close()

		go watchCheckinFiles(watcher, checkinDir)

		done := make(chan bool)
		err = watcher.Add(checkinDir)
		if err != nil {
			log.Fatal(err)
		}
		<-done
	}
}
