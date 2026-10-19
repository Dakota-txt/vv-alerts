package service

import (
	"errors"

	"gg-alerts/internal/models"
	"gg-alerts/internal/repository"
)

type DonationService struct {
	repository *repository.MemoryRepository
}

func NewDonationService(repository *repository.MemoryRepository) *DonationService {
	return &DonationService{
		repository: repository,
	}
}

func (s *DonationService) CreateDonation(
	streamerID int,
	amount int,
	message string,
) (models.Donation, error) {
	if amount <= 0 {
		return models.Donation{}, errors.New("сумма доната должна быть больше нуля")
	}

	_, err := s.repository.GetStreamer(streamerID)
	if err != nil {
		return models.Donation{}, err
	}

	donation := models.Donation{
		StreamerID: streamerID,
		Amount:     amount,
		Message:    message,
		Status:     "created",
	}

	savedDonation := s.repository.SaveDonation(donation)

	return savedDonation, nil
}

func (s *DonationService) GetDonations(streamerID int) ([]models.Donation, error) {
	_, err := s.repository.GetStreamer(streamerID)
	if err != nil {
		return nil, err
	}

	return s.repository.GetDonationsByStreamer(streamerID), nil
}

func (s *DonationService) GetTotalAmount(streamerID int) (int, error) {
	donations, err := s.GetDonations(streamerID)
	if err != nil {
		return 0, err
	}

	total := 0

	for _, donation := range donations {
		total += donation.Amount
	}

	return total, nil
}
