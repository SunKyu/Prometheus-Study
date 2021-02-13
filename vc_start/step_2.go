package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/VictoriaMetrics/metrics"
)

func main() {
	http.Handle("/", handle(home))
	http.HandleFunc("/metrics", handle(metricsPage))
	http.HandleFunc("/articles", handle(articles))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		metrics.GetOrCreateCounter(fmt.Sprintf(`requests_total{path=%q}`, r.URL.Path)).Inc()
		h.ServeHTTP(w, r)
	}
}

func metricsPage(w http.ResponseWriter, r *http.Request) {
	metrics.WritePrometheus(w, true)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!")
}

func articles(w http.ResponseWriter, r *http.Request) {
	metrics.GetOrCreateCounter(`requests_total{path="/articles"}`).Inc()
	resp, err := http.DefaultClient.Get("https://www.google.com/search?q=observability+articles")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprint(w, string(data))
}
