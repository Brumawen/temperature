package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Room holds the information about the room being monitored
type Room struct {
	Temperature float64   `json:"temp"`     // Room temperature
	Humidity    float64   `json:"humidity"` // Room humidity
	LastRead    time.Time `json:"lastread"` // Time the values were last read
}

// WriteTo serializes the entity and writes it to the http response
func (r *Room) WriteTo(w http.ResponseWriter) error {
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}
	w.Header().Set("content-type", "application/json")
	w.Write(b)
	return nil
}
