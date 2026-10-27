package httpapi

import (
	"encoding/json"
	"net/http"

	"gg-alerts/internal/service"
)

type Handler struct {
	donationService *service.DonationService
}

func NewHandler(donationService *service.DonationService) *Handler {
	return &Handler{
		donationService: donationService,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "ok",
		"service": "GG Alerts",
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}

func (h *Handler) ListDonations(w http.ResponseWriter, r *http.Request) {
	const streamerID = 1

	donations, err := h.donationService.GetDonations(streamerID)
	if err != nil {
		http.Error(w, "Не удалось получить донаты", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(donations); err != nil {
		return
	}
}
