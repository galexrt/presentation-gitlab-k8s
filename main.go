package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/handlers"
	"github.com/prometheus/common/version"
)

var ready = false

func main() {
	log.Info("Starting presentation-gitlab-k8s application..")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Write([]byte("Hello World!\n"))
		w.Write([]byte("Name: " + hostname + "\n"))
		w.Write([]byte("Version Info:\n"))
		w.Write([]byte(version.Print("app") + "\n"))
	})
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if ready {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500"))
		}
	})
	go func() {
		<-time.After(5 * time.Second)
		ready = true
		log.Info("Application is ready!")
	}()
	log.Info("Listen on :8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
