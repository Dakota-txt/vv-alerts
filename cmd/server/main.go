package main

import (
	"log"
	"net/http"

	"gg-alerts/internal/httpapi"
	"gg-alerts/internal/repository"
	"gg-alerts/internal/service"
)

func main() {
	repo := repository.NewMemoryRepository()
	donationService := service.NewDonationService(repo)
	handler := httpapi.NewHandler(donationService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /donations", handler.ListDonations)

	log.Println("GG Alerts запущен на порту 8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
