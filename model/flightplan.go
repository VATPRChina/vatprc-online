package model

type FlightPlan struct {
	Departure string `json:"departure"`
	Arrival string `json:"arrival"`
	Aircraft string `json:"aircraft_short"`
}

