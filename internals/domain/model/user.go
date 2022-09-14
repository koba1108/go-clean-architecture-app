package model

import "time"

type User struct {
	ID              int       `json:"id"`
	DisplayName     string    `json:"displayName"`
	FullName        string    `json:"fullName"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Birthday        time.Time `json:"birthday"`
	Age             int       `json:"age"`
	RegisteredAt    time.Time `json:"registeredAt"`
	LatestUpdatedAt time.Time `json:"latestUpdatedAt"`
}
