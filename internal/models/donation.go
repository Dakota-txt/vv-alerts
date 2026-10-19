package models

type Streamer struct {
	ID   int
	Name string
}

type Donation struct {
	ID         int
	StreamerID int
	Amount     int
	Message    string
	Status     string
}
