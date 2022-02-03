package controllers

import (
	"encoding/json"
	"fmt"
	"go-trading/app/models"
	"go-trading/config"
	"log"
	"net/http"
	"text/template"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request) {
	limit := 100
	duration := "1s"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)

	err := templates.ExecuteTemplate(w, "google.html", df.Candles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//jsonでエラーがあった時にjson形式で返す
type JSONerror struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func APIerror(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONerror{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
