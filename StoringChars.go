package asciiartweb

import (
	"bufio"
	"fmt"
	"os"
)


func StoringChars(font string) [][]string {
	var characters [][]string
	var currentCharacter []string
	/* storing the charecters as 2d arrays */
	var file *os.File
	var err error
	fmt.Print(font)
	switch font {
	case "standard":
		file, err = os.Open("../fonts/standard.txt")
	case "shadow":
		file, err = os.Open("../fonts/shadow.txt")
	case "thinkertoy":
		file, err = os.Open("../fonts/thinkertoy.txt")
	default:
		file, err = os.Open("../fonts/standard.txt")
	}
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
		// return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentCharacter) > 0 {
				characters = append(characters, currentCharacter)
				currentCharacter = nil
			}
		} else {
			currentCharacter = append(currentCharacter, line)
		}
	}
	/* appending the last charecter */
	if len(currentCharacter) > 0 {
		characters = append(characters, currentCharacter)
	}
	return characters
}
