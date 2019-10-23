package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	requestBody, err := json.Marshal(map[string]string{
		"name":  "oscar",
		"email": "none@gmail.com",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("before Post")
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(requestBody))
	fmt.Println("After Post")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

}
