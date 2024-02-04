package events

import "fmt"

const intro = `
I need a list of event details formatted into a CSV. The CSV should have the following columns with headers: Name, RecursOn, Date, StartTime, EndTime, Description, Price. Please adhere strictly to the formats and rules outlined below for each column:

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

Given the list of event details:
`

const outro = `
Please create a CSV with the details as per the instructions above. Make sure to check that each piece of data is in the correct column and that the CSV is free from formatting errors or extraneous text.
`

func getPrompt(pageText string) string {
	return fmt.Sprintf(
		"%s\n"+
			"%s\n\n"+
			"%s",
		intro, pageText, outro)
}
