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

// @Summary Update Harian Indonesia
// @Description An API that used to get data of Indonesian daily covid case
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /update-harian-indonesia [get]
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

// @Summary Pemeriksaan dan Vaksinasi
// @Description An API that used to get data of Indonesian covid checkup and vaccination
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /pemeriksaan-dan-vaksinasi [get]
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

// @Summary KasusSeluruhProvinsi
// @Description An API that used to get data of Indonesian covid based on provinces
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /kasus-seluruh-provinsi [get]
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

// @Summary Risk Score Kecamatan
// @Description An API that used to get Indonesia risk score data based on districs
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /risk-score-kecamatan [get]
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

// @Summary Rumah Sakit Rujukan
// @Description An API used to get Indonesian hospital that treats covid patients data
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /rumah-sakit-rujukan [get]
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

// @Summary Labolatorium
// @Description An API used to get Indonesian labolatorium that treats covid patients data
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /labolatorium-rujukan [get]
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

// @Summary Data Harian Provinsi
// @Description An API used to get Indonesian daily covid case update based on provinces
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /data-harian-provinsi [get]
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

// @Summary Risk Score Provinsi
// @Description An API that used to get Indonesia risk score data based on provinces
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /risk-score-provinsi [get]
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

// @Summary Berita Covid
// @Description An API that used to get Indonesia news regarding covid
// @Tags v1
// @Produce json
// @Success 200 {object} swaggerresponse.SwaggerSuccessResponse
// @Failure 500 {object} swaggerresponse.SwaggerErrorResponse
// @Router /berita-covid [get]
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
	http.HandleFunc("/v1/update-harian-indonesia", UpdateHarianIndonesia)
	http.HandleFunc("/v1/pemeriksaan-dan-vaksinasi", PemeriksaanDanVaksinasi)
	http.HandleFunc("/v1/kasus-seluruh-provinsi", KasusSeluruhProvinsi)
	http.HandleFunc("/v1/risk-score-kecamatan", RiskScoreKecamatan)
	http.HandleFunc("/v1/rumah-sakit-rujukan", RumahSakitRujukan)
	http.HandleFunc("/v1/labolatorium-rujukan", LabolatoriumRujukan)
	http.HandleFunc("/v1/data-harian-provinsi", DataHarianProvinsi)
	http.HandleFunc("/v1/risk-score-provinsi", RiskScoreProvinsi)
	http.HandleFunc("/v1/berita-covid", BeritaCovid)

	log.Print("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
