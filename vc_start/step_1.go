package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/articles", articles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func articles(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Get("https://www.google.com/search?q=observability+articles")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}