package events

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/nazmalik/edinburgh-events/chatgpt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// 700 gets to event #11 of 23
	maxTokens          = 4096
	csvExampleFilename = "events.csv"
)

// Fetch takes the text of a page and hits the ChatGPT API to extract event data
func Fetch(pageText string, venueName string, url string) ([]Event, error) {
	preprocessingRsp, err := chatgpt.SendPrompt(getPreprocessingPrompt(pageText), maxTokens)
	if err != nil {
		return nil, err
	}

	jsonRsp, err := chatgpt.SendPrompt(getJSONPrompt(preprocessingRsp.FirstChoiceOutput()), maxTokens)
	if err != nil {
		return nil, err
	}

	fmt.Println(jsonRsp.FirstChoiceOutput())

	events, err := getEventsFromJSON(jsonRsp.FirstChoiceOutput())
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON content: %v", err)
	}

	enrichedEvents := make([]Event, len(events))
	for i, event := range events {
		event.Venue = venueName
		event.URL = url
		enrichedEvents[i] = event
	}

	// Convert the events to JSON
	jsonData, err := json.Marshal(enrichedEvents)
	if err != nil {
		return nil, fmt.Errorf("error converting events to JSON: %v", err)
	}

	jsonFile := os.Getenv("FRONTEND_EVENTS_JSON_PATH")
	if err := os.WriteFile(jsonFile, jsonData, 0644); err != nil {
		log.Fatalf("Error writing events JSON file: %v", err)
	}

	return enrichedEvents, nil

}

// getEventsFromCSV takes a CSV-formatted string and returns an array of Event structs
func getEventsFromCSV(csvString string, venueName string, url string) ([]Event, error) {
	reader := csv.NewReader(strings.NewReader(csvString))
	var events []Event
	isHeaderRecord := true

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if isHeaderRecord {
			isHeaderRecord = false
			continue
		}

		event := Event{
			Name:        record[0],
			RecursOn:    record[1],
			Date:        record[2],
			StartTime:   record[3],
			EndTime:     record[4],
			Description: record[5],
			Price:       record[6],
		}

		events = append(events, event)
	}

	return events, nil
}
