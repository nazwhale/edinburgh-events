package events

type Event struct {
	Name         string // e.g. "Live Music Night"
	RecursOn     string // e.g. "Every Tuesday". Should be empty if event does not recur
	Date         string // e.g. "2024-05-14". Empty if event recurs. If unsure of year, use 2024. One of Date or RecursOn must be present
	StartTime    string // e.g. "14:00". 24hr time format as a string. e.g. 14:00 for 2pm
	EndTime      string // e.g. "18:00". 24hr time format as a string, may be empty. e.g. 18:00 for 6pm.
	Description  string // e.g. "Jazz & Blues". Summarise type of event in less than 6 words. Not including price (£) or time information.
	Price        int    // in pennies, e.g., 1000 for £10; £8otd is 800
	UnknownPrice bool   // If true, the price is unknown
}
