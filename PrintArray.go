package asciiartweb

// import (
// 	"strings"
// )


func PrintArray(text string, banner string) string {
	var PrintThis string
	characters := StoringChars(banner)
	PrintThis += "\n"
	var ascii rune
	//newLine := strings.Split(text, " ")
	//for _, v := range newLine {
		/* printing the charecters */
		for i := 0; i < 8; i++ {
			for _, char := range text {
				for range characters[ascii][0:8] {
					ascii = char - 32
					PrintThis += characters[ascii][0:8][i]

					break
				}
			}
			PrintThis += "\n"
		}
	//}
	return PrintThis
}
