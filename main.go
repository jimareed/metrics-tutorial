package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ItemList : list of items
type ItemList []struct {
	Item string `json:"item"`
}

func main() {
	http.HandleFunc("/items", items)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)
}

func items(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[{\"item\":\"apple\"}, {\"item\":\"orange\"}, {\"item\":\"pear\"}]\n")
}

func test(w http.ResponseWriter, r *http.Request) {

	if count() == 3 {
		io.WriteString(w, "{\"success\":\"")
		io.WriteString(w, "true")
		io.WriteString(w, "\"}\n")
	} else {
		io.WriteString(w, "{\"success\":\"")
		io.WriteString(w, "false")
		io.WriteString(w, "\"}\n")
	}

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
