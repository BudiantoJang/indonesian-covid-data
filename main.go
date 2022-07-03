package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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
	http.HandleFunc("/update-harian-indonesia", UpdateHarianIndonesia)

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
