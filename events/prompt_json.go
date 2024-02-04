package events

import "fmt"

const jsonIntro = `
I need a list of event details formatted into JSON. The JSON should have the following shape: 
{
	name, // string. e.g. "Live Music Night"
	recurs_on,  // string. e.g. "Every Tuesday". Should be "" if event does not recur
    date, // string.e.g. "2024-05-14". Should be "" if event recurs. If unsure of year, use 2024. One of date or recurs_on must be present
	start_time,  // string. e.g. 2pm is "14:00". 24hr time format as a string. 
	end_time, // string. e.g. 6pm is "18:00". 24hr time format as a string, may be empty. 
	description, // string. e.g. "Jazz & Blues". Summarise type of event in less than 6 words. Not including price (£) or time information.
	price_in_pennies, // string. in pennies, e.g., 1000 for £10; £8otd is 800. Free is 0. Leave as "" if price is unclear.
	is_price_unknown, // boolean, true if price is unknown or unconfirmed or tbc
}

Note: 
- For events with multiple dates, each date should be a separate record with the same "name"
- Output should begin with a square bracket "[" and end with a square bracket "]"

Please extract the event details from the text and output an array of JSON objects, following the rules above. 
Example output for 2 events: 
[
    {
        "name": "Pub Quiz with John Doe",
		"recurs_on": "Every Tuesday",
		"date": "",
		"start_time": "20:00",
		"end_time": "22:00",
		"description": "Quiz",
		"price_in_pennies": "500",
		"is_price_unknown": false,
	},
	{
		"name": "Pub Quiz with John Doe",	
		"recurs_on": "Every Tuesday",
		"date": "",
		"start_time": "20:00",
		"end_time": "22:00",
		"description": "Quiz",
		"price_in_pennies": "500",
		"is_price_unknown": false,
	}
]

It is very important that there is a JSON object for every single event below.

Events:
`

const jsonOutro = ``

func getJSONPrompt(pageText string) string {
	return fmt.Sprintf(
		"%s\n"+
			"%s\n\n"+
			"%s",
		jsonIntro, pageText, jsonOutro)
}
