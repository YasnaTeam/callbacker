package server

import (
	"github.com/sirupsen/logrus"
	"strconv"
	"net/http"
	"fmt"
)

var log *logrus.Logger

func Initialize(port uint, logger *logrus.Logger) {
	log = logger
	log.Info("Server started on port " + strconv.Itoa(int(port)) + "...")
	http.HandleFunc("/")
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(int(port)), nil))
}


func hello(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	log.WithRequest
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}