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

var ready = false
var addr = flag.String("listen-address", ":8000", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	log.Info("Starting presentation-gitlab-k8s application..")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Write([]byte("Hello Golang DevOpsCon Conference 2019 Berlin!\n"))
		w.Write([]byte("Hostname: " + hostname + "\n"))
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
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		<-time.After(5 * time.Second)
		ready = true
		log.Info("Application is ready!")
	}()
	log.Info("Listen on " + *addr)
	log.Fatal(http.ListenAndServe(*addr, handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}
