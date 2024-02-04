package events

import "fmt"

const correctingIntro = `
Given the .csv:
`

const correctingOutro = `
Please correct the .csv to the best of your ability. Refer to the list of events below. Some may be missing from the .csv, some may be incomplete, some may be perfect already.
Make sure to check that each piece of data is in the correct column and that the CSV is free from formatting errors.

.csv header: Name,RecursOn,Date,StartTime,EndTime,Description,Price

Here's how a correctly formatted CSV row would look for a recurring event:
"Pub Quiz with John Doe", "Every Tuesday", "", "20:00", "22:00", "Trivia Night", "500"

And for a one-time event:
"New Year's Eve Gala", "", "2024-12-31", "20:00", "01:00", "Gala Event", "7500"

For each column, adhere to the following rules:

1. Name: Use the exact title of the event.
   - Example: "Live Music Night"
   
2. RecursOn: Indicate the day the event regularly occurs, leave blank if it's a one-time event.
   - Example: "Every Tuesday"
   - Example for one-time event: ""
   
3. Date: Provide the specific date in YYYY-MM-DD format, leave blank if the event recurs. Assume 2024 if year is not clear.
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
   
6. Description: A concise summary of the event. 
   - Example: "Jazz & Blues", "Stand-up Comedy", "Folk Music", "Art Workshop", "Swing Dancing Social"
   
7. Price: The numeric value in pennies, '0' if free. Do not include currency symbols.
   - Example: "1000" for £10
   - Example for a free event: "0"

Notes: 
- If time or date information is "TBC" or unconfirmed, leave blank

Input that was used to generate the .csv:

`

func getCorrectingPrompt(pageText string, eventsList string) string {
	return fmt.Sprintf(
		"%s\n"+
			"%s\n\n"+
			"%s"+
			"%s",
		correctingIntro, pageText, correctingOutro, eventsList)
}
