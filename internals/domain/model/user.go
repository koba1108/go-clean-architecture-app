package model

import (
	"fmt"
	"time"
)

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

func NewUser(displayName, firstName, lastName string, birthday time.Time) (*User, error) {
	if displayName == "" {
		return nil, fmt.Errorf("displayName is required")
	}
	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if lastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	if birthday.IsZero() {
		return nil, fmt.Errorf("birthday is required")
	}
	age := int(time.Now().Sub(birthday).Hours() / 24 / 365)
	if age < 20 {
		return nil, fmt.Errorf("age must be over 20")
	}
	return &User{
		DisplayName:     displayName,
		FirstName:       firstName,
		LastName:        lastName,
		FullName:        fmt.Sprintf("%s %s", lastName, firstName),
		Birthday:        birthday,
		Age:             age,
		RegisteredAt:    time.Now(),
		LatestUpdatedAt: time.Now(),
	}, nil
}
