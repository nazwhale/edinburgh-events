package events

import (
	"encoding/csv"
	"fmt"
	"github.com/nazmalik/edinburgh-events/chatgpt"
	"io"
	"strconv"
	"strings"
)

const (
	// 700 gets to event #11 of 23
	maxTokens          = 4000
	csvExampleFilename = "events.csv"
)

// Fetch takes the text of a page and hits the ChatGPT API to extract event data
func Fetch(pageText string) error {
	preprocessingRsp, err := chatgpt.SendPrompt(getPreprocessingPrompt(pageText), maxTokens)
	if err != nil {
		return err
	}

	rsp, err := chatgpt.SendPrompt(getPrompt(preprocessingRsp.FirstChoiceOutput()), maxTokens)
	if err != nil {
		return err
	}

	csvString := rsp.FirstChoiceOutput()

	if err := writeToCSV(csvString, csvExampleFilename); err != nil {
		return fmt.Errorf("error writing to CSV file: %v", err)
	}

	fmt.Println("Written .csv 🎉")

	events, err := parseCSVFromString(csvString)
	if err != nil {
		return fmt.Errorf("error parsing CSV content: %v", err)
	}

	fmt.Println("Parsed events 🎉")

	// Print the parsed events to verify
	for _, event := range events {
		fmt.Printf("%+v\n", event)
	}

	return nil
}

// parseCSVFromString takes a CSV-formatted string and returns an array of Event structs
func parseCSVFromString(csvString string) ([]Event, error) {
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

		price, err := strconv.Atoi(record[6])
		if err != nil {
			fmt.Printf("Error converting price for event %s: %v\n", record[0], err)
			continue // Skip this record
		}

		event := Event{
			Name:        record[0],
			RecursOn:    record[1],
			Date:        record[2],
			StartTime:   record[3],
			EndTime:     record[4],
			Description: record[5],
			Price:       price,
		}

		events = append(events, event)
	}

	return events, nil
}
