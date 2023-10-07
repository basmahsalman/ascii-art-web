package asciiartweb

import (
	"fmt"
	//"os"
)

func ValidatingInput(text string) {
	/* cheking if the cherecter is valid */
	for i := 0; i < len(text); i++ {
		if text[i] < 32 || text[i] > 126 {
			fmt.Println("invalid charecter")
			//os.Exit(0)
		}
	}
}
