package asciiartweb

import (
	//"fmt"
	"strings"
)

var PrintThis string

func PrintArray(text string, banner string) string {
	characters := StoringChars(banner)
	var ascii rune
	newLine := strings.Split(text, "\n")
	for _, v := range newLine {
		// if v == "" {
		// 	fmt.Println()
		// 	continue
		// }
		PrintThis = ""
		/* printing the charecters */
		for i := 0; i < 8; i++ {
			for _, char := range v {
				for range characters[ascii][0:8] {
					ascii = char - 32
					PrintThis += characters[ascii][0:8][i]

					break
				}
			}
			PrintThis += "\n"
		}
	}
	return PrintThis
}
