package main

import (
	"log"
	"net/http"
	"weather_api/internal"
)

func main() {
	logger, err := internal.NewLogger("server.log")
	if err != nil {
		log.Fatal(err)
	}

	handler := internal.NewWeatherHandler(logger)

	router := http.NewServeMux()
	router.HandleFunc("/", handler.HomeHandler)
	router.HandleFunc("/weather", handler.WeatherHandler)

	fileServer := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fileServer))

	logger.Println("Сервер запущен на порту 80")
	log.Fatal(http.ListenAndServe(":80", router))
}
