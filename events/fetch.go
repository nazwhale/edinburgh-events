package events

import (
	"encoding/csv"
	"fmt"
	"github.com/nazmalik/edinburgh-events/chatgpt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	// 700 gets to event #11 of 23
	maxTokens = 2000
)

// Fetch takes the text of a page and hits the ChatGPT API to extract event data
func Fetch(pageText string) error {
	promptIntro := `I need the event details formatted into a CSV. The CSV should have the following columns with headers: Name, RecursOn, Date, StartTime, EndTime, Description, Price. Please adhere strictly to the formats and rules outlined below for each column:

For each column, adhere to the following examples:

1. Name: Use the exact title of the event.
   - Example: "Live Music Night"
   
2. RecursOn: Indicate the day the event regularly occurs, leave blank if it's a one-time event.
   - Example: "Every Tuesday"
   - Example for one-time event: ""
   
3. Date: Provide the specific date in YYYY-MM-DD format, leave blank if the event recurs.
   - Example for one-time event: "2024-01-21" for Sun 21st Jan
   - Example for a recurring event: ""
   - Invalid: "Wed 10th Dec" (correct format: "2024-12-10")
   
4. StartTime: Indicate the starting time in 24-hour HH:MM format.
   - Example: "19:00" for 7pm
   - Example: "14:00" for 2pm-4pm (EndTime would be "16:00")
   - Example if not specified: ""

   
5. EndTime: Provide the ending time in 24-hour HH:MM format, leave blank if not specified.
   - Example: "21:00"
   - Example if not specified: ""
   
6. Description: A concise summary of the event, not including price or time.
   - Example: "Jazz & Blues"
   - Example: "Folk Music"
   
7. Price: The numeric value in pennies, '0' if free. Do not include currency symbols.
   - Example: "1000" for £10
   - Example for a free event: "0"

Here's how a correctly formatted CSV row would look for a recurring event:

"Pub Quiz with John Doe", "Every Tuesday", "", "20:00", "22:00", "Trivia Night", "500"

And for a one-time event:

"New Year's Eve Gala", "", "2024-12-31", "20:00", "01:00", "Gala Event", "7500"

Notes: 
- If a field's data is ambiguous or doesn't fit the specified format, exclude it from the output.
- One of Date or RecursOn must be present
- For events with multiple dates, each date should be a separate record with the same Name and Description.
- Ensure that each data point is placed in the appropriate column.

Please extract the event details from the text and input them into the correct columns, following the rules above. If the provided information does not fit the format or is ambiguous, omit it from the CSV. 

`

	promptOutroCSV := "Please create a CSV with the details as per the instructions above. Make sure to check that each piece of data is in the correct column and that the CSV is free from formatting errors or extraneous text."

	prompt := fmt.Sprintf(
		"%s\n"+
			"Given the event page content:\n\n"+
			"%s\n\n"+
			"%s",
		promptIntro, pageText, promptOutroCSV)

	fmt.Println(prompt)

	fmt.Println("\nSending prompt to ChatGPT API 🚀")

	rsp, err := chatgpt.SendPrompt(prompt, maxTokens)
	if err != nil {
		return err
	}

	// Print the response from the API
	for _, c := range rsp.Choices {
		fmt.Println()
		fmt.Println(c.Message.Content)

		if c.FinishReason == chatgpt.FinishReasonLength {
			fmt.Println("Response truncated due to length 🐍")
			fmt.Println()
		}
	}

	// Print usage info
	fmt.Printf("\n🇹🇰\nTokens used: %d\nPrompt: %d\nCompletion: %d\n\n", rsp.Usage.TotalTokens, rsp.Usage.PromptTokens, rsp.Usage.CompletionTokens)

	csvString := rsp.Choices[0].Message.Content

	if err := writeToCSV(csvString, "events.csv"); err != nil {
		return fmt.Errorf("error writing to CSV file: %v", err)
	}

	fmt.Println("Written .csv 🎉")

	// Convert the CSV string to a reader
	reader := csv.NewReader(strings.NewReader(csvString))

	var events []Event

	isHeaderRecord := true

	for {

		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// End of file reached, break out of the loop
				break
			}

			return fmt.Errorf("error reading CSV record: %v", err)
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
		}

		priceRecord := record[6]
		if priceRecord == "" {
			event.UnknownPrice = true
		} else {
			price, err := strconv.Atoi(priceRecord)
			if err != nil {
				fmt.Printf("Error converting price for event %s: %v\n", record[0], err)
				continue // Skip this record
			}

			event.Price = price
		}

		events = append(events, event)
	}

	fmt.Println("Parsed events 🎉")

	// Print the parsed events to verify
	for _, event := range events {
		fmt.Printf("%+v\n", event)
	}

	return nil
}

// writeToCSV takes CSV content as a string and writes it to a specified file.
// It returns an error if the file cannot be created or written to.
func writeToCSV(csvContent, fileName string) error {
	// Create or overwrite the file
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Write the CSV content to the file
	_, err = file.WriteString(csvContent)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil // No error occurred
}
