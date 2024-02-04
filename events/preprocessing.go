package events

import "fmt"

const preprocessingIntro = `
Given the event page content:
`

const preprocessingOutro = `
Pull out details of each individual event into a numbered list.
Do not include any other text or formatting.
`

func getPreprocessingPrompt(pageText string) string {
	return fmt.Sprintf(
		"%s\n"+
			"%s\n\n"+
			"%s",
		preprocessingIntro, pageText, preprocessingOutro)
}
