package events

import "fmt"

const preprocessingIntro = `
Given the event page content:
`

const preprocessingOutro = `
Pull out details of each individual event into a numbered list.
Include price, time, and date information where available.
Do not include any other text or formatting.
It is very important that no events are left out.

Example:
1. Pub Quiz - Every Tuesday at 8pm, FREE
2. Pictureskew + Dear Srrrz + Alexx Munro + Sideline Burnout (Synth Punk) - Sat 6th Jan, Doors 7:30 pm, £5
3. Drink & Draw - Sun 7th Jan, 2pm-4pm, £10.50
4. Ableton User Group (Workshop/Talk/Ableton) - Mon 22nd Jan, Time TBC, FREE
`

func getPreprocessingPrompt(pageText string) string {
	return fmt.Sprintf(
		"%s\n"+
			"%s\n"+
			"%s",
		preprocessingIntro, pageText, preprocessingOutro)
}
