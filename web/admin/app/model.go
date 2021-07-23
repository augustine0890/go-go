package main

import "time"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	StartDate time.Time
	Position  string
	TotalPTO  float64
	Status    string
	TimesOff  []TimeOff
}

type TimeOff struct {
	Type      string
	Amount    float64
	StartDate time.Time
	Status    string
}

var employees = map[string]Employee{
	"962134": {
		ID:        962134,
		FirstName: "Jennifer",
		LastName:  "Watson",
		Position:  "CEO",
		StartDate: time.Now().Add(-12 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  30,
	},
	"176158": {
		ID:        176158,
		FirstName: "Allison",
		LastName:  "Jane",
		Position:  "COO",
		StartDate: time.Now().Add(-4 * time.Hour * 24 * 365),
		Status:    "Active",
		TotalPTO:  20,
	},
	"160898": {
		ID:        160898,
		FirstName: "Aakar",
		LastName:  "Uppal",
		Position:  "CTO",
		StartDate: time.Now().Add(-6 * time.Hour * 24 * 365),
		TotalPTO:  20,
	},
	"297365": {
		ID:        297365,
		FirstName: "Jonathon",
		LastName:  "Anderson",
		StartDate: time.Now().Add(-11 * time.Hour * 24 * 365),
		Position:  "Worker Bee",
		TotalPTO:  30,
	},
}

var TimesOff = map[string][]TimeOff{
	"962134": {
		{
			Type:      "Holiday",
			Amount:    8.,
			StartDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			Status:    "Taken",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2021, 8, 16, 0, 0, 0, 0, time.UTC),
			Status:    "Scheduled",
		},
		{
			Type:      "PTO",
			Amount:    16.,
			StartDate: time.Date(2021, 12, 8, 0, 0, 0, 0, time.UTC),
			Status:    "Requested",
		},
	},
}
