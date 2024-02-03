package html

import (
	"bytes"
	"golang.org/x/net/html"
	"strings"
)

// ToText takes an HTML string and returns its text content without tags.
func ToText(htmlContent string) (string, error) {
	// Parse the HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return "", err
	}

	// Buffer to accumulate text nodes
	var buf bytes.Buffer

	// Extract text nodes into the buffer
	extractTextNodes(doc, &buf, false)

	// Further process the buffer to remove any remaining excessive whitespace.
	// This could include multiple spaces reduced to a single space and trimming the result.
	result := strings.Join(strings.Fields(buf.String()), " ")

	// Remove common useless phrases
	result = removePhrases(result)

	// Return the accumulated text
	return result, nil
}

// removePhrases takes an input text and a slice of phrases to remove from the text.
// It returns the text with all specified phrases removed.
func removePhrases(inputText string) string {
	// phrasesToRemove should be longer than 1 word, so that we don't replace common words by mistake
	var phrasesToRemove = []string{
		"book now",
		"get tickets",
		"reserve a space",
	}

	// Iterate over each phrase to remove
	for _, phrase := range phrasesToRemove {
		// Generate common casings of the phrase
		lowerPhrase := strings.ToLower(phrase)
		upperPhrase := strings.ToUpper(phrase)
		titlePhrase := strings.Title(phrase)

		// Attempt to remove each casing of the phrase
		inputText = strings.ReplaceAll(inputText, lowerPhrase, "")
		inputText = strings.ReplaceAll(inputText, upperPhrase, "")
		inputText = strings.ReplaceAll(inputText, titlePhrase, "")
	}

	return inputText
}

// extractTextNodes traverses the HTML node tree and accumulates text nodes' data.
func extractTextNodes(n *html.Node, buf *bytes.Buffer, skip bool) {
	// If the current node is a script or style tag, set skip to true
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		skip = true
	}
	if n.Type == html.TextNode && !skip {
		// Add a space before the text to ensure separation from previous text
		buf.WriteString(" ")
		buf.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractTextNodes(c, buf, skip)
	}
}
