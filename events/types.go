package events

import (
	"encoding/json"
	"fmt"
)

type Event struct {
	Name        string `json:"name"`             // e.g. "Live Music Night"
	RecursOn    string `json:"recurs_on"`        // e.g. "Every Tuesday". Should be empty if event does not recur
	Date        string `json:"date"`             // e.g. "2024-05-14". Empty if event recurs. If unsure of year, use 2024. One of Date or RecursOn must be present
	StartTime   string `json:"start_time"`       // e.g. "14:00". 24hr time format as a string. e.g. 14:00 for 2pm
	EndTime     string `json:"end_time"`         // e.g. "18:00". 24hr time format as a string, may be empty. e.g. 18:00 for 6pm.
	Description string `json:"description"`      // e.g. "Jazz & Blues". Summarise type of event in less than 6 words. Not including price (£) or time information.
	Price       string `json:"price_in_pennies"` // in pennies, e.g., 1000 for £10; £8otd is 800
	Venue       string `json:"venue"`            // e.g. "The Old Inn"
	URL         string `json:"url"`              // e.g. "https://www.theoldinn.com/live-music-night"
}

// getEventsFromJSON takes a JSON-formatted string and returns an array of Event structs
func getEventsFromJSON(jsonString string) ([]Event, error) {
	var events []Event
	if err := json.Unmarshal([]byte(jsonString), &events); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	// Optional: Here you can handle any post-processing if necessary
	// For example, converting certain fields, handling defaults, etc.

	return events, nil
}
