package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(w http.ResponseWriter, r *http.Request) {
	url := "https://dummyjson.com/products"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Fprint(w, responseString)
}

func main() {
	http.HandleFunc("/", Fetch)
	if err := http.ListenAndServe(":3000", nil); err != nil {
	}
}
