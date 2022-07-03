package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func UpdateHarianIndonesia(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://api-covid-indonesia.herokuapp.com/UpdateHarianIndonesia")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result := string(body)
	w.Write([]byte(result))
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/update-harian-indonesia", UpdateHarianIndonesia)

	log.Print("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
