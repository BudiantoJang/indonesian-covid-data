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

func PemeriksaanDanVaksinasi(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/pemeriksaan-vaksinasi.json")
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

func KasusSeluruhProvinsi(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/prov.json")
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

func RiskScoreKecamatan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/kecamatan_rawan.json")
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

func RumahSakitRujukan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/lab.json")
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

func LabolatoriumRujukan(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/rs.json")
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

func DataHarianProvinsi(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/prov_time.json")
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

func RiskScoreProvinsi(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://data.covid19.go.id/public/api/skor.json")
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

func BeritaCovid(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://newsapi.org/v2/everything?q=covid&apiKey=4f8b1414b1634e198f8758eef4d8daf2&domains=detik.com")
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

// @title Indonesian Covid Data API Documentation
// @description This is an API documentation for Indonesian Covid Data API that is generated using OpenAPI 2.0 specification
// @version v1
// @BasePath /v1
func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("v1/update-harian-indonesia", UpdateHarianIndonesia)
	http.HandleFunc("v1/pemeriksaan-dan-vaksinasi", PemeriksaanDanVaksinasi)
	http.HandleFunc("v1/kasus-seluruh-provinsi", KasusSeluruhProvinsi)
	http.HandleFunc("v1/risk-score-kecamatan", RiskScoreKecamatan)
	http.HandleFunc("v1/rumah-sakit-rujukan", RumahSakitRujukan)
	http.HandleFunc("v1/labolatorium-rujukan", LabolatoriumRujukan)
	http.HandleFunc("v1/data-harian-provinsi", DataHarianProvinsi)
	http.HandleFunc("v1/risk-score-provinsi", RiskScoreProvinsi)
	http.HandleFunc("v1/berita-covid", BeritaCovid)

	log.Print("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
