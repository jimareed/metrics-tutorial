package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// ItemList : list of items
type ItemList []struct {
	Item string `json:"item"`
}

func main() {
	http.HandleFunc("/items", items)
	http.HandleFunc("/test", test)
	http.HandleFunc("/", health)
	http.ListenAndServe(":8080", nil)
}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}

func items(w http.ResponseWriter, r *http.Request) {
	responseTime, err := strconv.Atoi(os.Getenv("RESPONSE_TIME"))
	if err != nil {
		responseTime = 50
	}

	failureRate, err := strconv.Atoi(os.Getenv("FAILURE_RATE"))
	if err != nil {
		failureRate = 10
	}

	random := rand.Intn(responseTime * 2)
	delay := time.Millisecond * time.Duration(random)

	time.Sleep(delay)

	if rand.Intn(100) >= failureRate {
		io.WriteString(w, "[{\"item\":\"apple\"}, {\"item\":\"orange\"}, {\"item\":\"pear\"}]\n")
	} else {
		io.WriteString(w, "[{\"error\":\"true\"}]\n")
	}

}

func test(w http.ResponseWriter, r *http.Request) {

	iterations, err := strconv.Atoi(os.Getenv("ITERATIONS"))
	if err != nil {
		iterations = 100
	}

	success := 0
	fail := 0

	start := time.Now()

	for i := 0; i < iterations; i++ {
		if count() == 3 {
			success++
		} else {
			fail++
		}
	}

	t := time.Now()
	elapsed := t.Sub(start)

	s := strconv.Itoa(success)
	f := strconv.Itoa(fail)

	io.WriteString(w, "{\"success\":\"")
	io.WriteString(w, s)
	io.WriteString(w, "\" , \"fail\":\"")
	io.WriteString(w, f)
	io.WriteString(w, "\" , \"elapsed\":\"")
	io.WriteString(w, elapsed.String())
	io.WriteString(w, "\"}\n")

}

func count() int {

	url := os.Getenv("ITEMS_SERVICE_URL") + "/items"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return -1
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1
	}

	var i ItemList

	err = json.Unmarshal(body, &i)
	if err != nil {
		return -1
	}

	return len(i)
}
