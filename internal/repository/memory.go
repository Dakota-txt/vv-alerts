package repository

import (
	"errors"

	"gg-alerts/internal/models"
)

type MemoryRepository struct {
	streamers map[int]models.Streamer
	donations []models.Donation
	nextID    int
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		streamers: make(map[int]models.Streamer),
		donations: make([]models.Donation, 0),
		nextID:    1,
	}
}

func (r *MemoryRepository) AddStreamer(streamer models.Streamer) {
	r.streamers[streamer.ID] = streamer
}

func (r *MemoryRepository) GetStreamer(id int) (models.Streamer, error) {
	streamer, exists := r.streamers[id]
	if !exists {
		return models.Streamer{}, errors.New("стример не найден")
	}

	return streamer, nil
}

func (r *MemoryRepository) SaveDonation(donation models.Donation) models.Donation {
	donation.ID = r.nextID
	r.nextID++

	r.donations = append(r.donations, donation)

	return donation
}

func (r *MemoryRepository) GetDonationsByStreamer(streamerID int) []models.Donation {
	donations := make([]models.Donation, 0)

	for _, donation := range r.donations {
		if donation.StreamerID == streamerID {
			donations = append(donations, donation)
		}
	}

	return donations
}
