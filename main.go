package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	// -----------------------------------------------
	// add prometheus go client import:
	//
	// "github.com/prometheus/client_golang/prometheus/promhttp"
	//
	// -----------------------------------------------
)

// ItemList : list of items
type ItemList []struct {
	Item string `json:"item"`
}

func main() {
	http.HandleFunc("/items", items)
	http.HandleFunc("/", health)

	// -----------------------------------------------
	// add metrics handler:
	//
	// http.Handle("/metrics", promhttp.Handler())
	//
	// -----------------------------------------------

	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}

func items(w http.ResponseWriter, r *http.Request) {

	failureRate, err := strconv.Atoi(os.Getenv("FAILURE_RATE"))
	if err != nil {
		failureRate = 0
	}

	if rand.Intn(100) >= failureRate {
		io.WriteString(w, "[{\"item\":\"apple\"}, {\"item\":\"orange\"}, {\"item\":\"pear\"}]\n")
	} else {
		http.Error(w, http.StatusText(500), 500)
	}

}
