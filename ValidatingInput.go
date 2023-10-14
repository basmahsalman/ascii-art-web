package asciiartweb

import (
	"fmt"
)

func ValidatingInput(text string) string {
	/* cheking if the cherecter is valid */
	for i := 0; i < len(text); i++ {
		// text[i]=10 && 13 is related to the 'Enter'/New Line
		if (text[i] < 32 || text[i] > 126) && text[i]!=10 && text[i]!=13 {
			fmt.Println("invalid charecter")
			text = ""
		}
	}
	return text
}
