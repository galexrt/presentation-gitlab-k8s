/*
MIT License

Copyright (c) 2018 Alexander Trost

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
)

var (
	ready = false
	addr  = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")
)

func main() {
	flag.Parse()
	log.Info("Starting web demo app application..")

	// Register ("/metrics") endpoint for Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	// Index ("/") handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		// When unable to get the hostname, just set it to `N/A`
		// remember this is just a demo app ;-)
		if err != nil {
			hostname = "N/A"
		}
		w.Write([]byte("Hello!\n"))
		w.Write([]byte("Hostname: " + hostname + "\n"))
		w.Write([]byte("Version Info:\n"))
		w.Write([]byte(version.Print("app") + "\n"))
	})

	// Demo "/health" handler
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
		// Demo purpose 7 second delay till application "/health" endpoint
		// will be ready
		<-time.After(7 * time.Second)
		ready = true
		log.Info("Application is ready!")
	}()

	// Start the web server
	log.Info("Listen on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
