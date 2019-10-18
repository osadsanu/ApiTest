package main

//version of get using Golang - Platzi (API REST course)

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "https://xkcd.com/info.0.json"
	client := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	obj := make(map[string]interface{})
	jsonErr := json.Unmarshal(body, &obj)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//jsonData := string(body)
	fmt.Println(obj["img"])
}
