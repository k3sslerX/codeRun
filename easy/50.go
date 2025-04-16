package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var textStr string
	for scanner.Scan() {
		textStr += " " + scanner.Text()
	}
	//textStr := scanner.Text()
	textRune := []rune(textStr)
	dict := make(map[string]int)
	var curWord string
	for i := 0; i < len(textRune); i++ {
		if textStr[i] == ' ' {
			dict[curWord]++
			curWord = ""
		} else {
			curWord = curWord + string(textRune[i])
		}
	}
	max := 0
	maxWord := ""
	for k, v := range dict {
		if v > max {
			max = v
			maxWord = k
		} else if v == max {
			if k < maxWord {
				maxWord = k
			}
		}
	}
	fmt.Println(maxWord)
}
