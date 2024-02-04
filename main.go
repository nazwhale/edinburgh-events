package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nazmalik/edinburgh-events/events"
	"github.com/nazmalik/edinburgh-events/html"
	"log"
)

func main() {
	url := "https://leithdepot.com/events.html"
	venueName := "Leith Depot"

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Fetch the HTML content of the webpage
	htmlContent, err := html.FetchHTML(url)
	if err != nil {
		fmt.Println("Error fetching HTML content:", err)
		return
	}

	// Convert HTML to plain text
	text, err := html.ToText(htmlContent)
	if err != nil {
		fmt.Println("Error converting HTML to text:", err)
		return
	}

	fetchedEvents, err := events.Fetch(text, venueName, url)
	if err != nil {
		fmt.Println("Error fetching events:", err)
		return
	}

	fmt.Println(fetchedEvents)
}
