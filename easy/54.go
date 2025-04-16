package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	dict := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		m, _ := strconv.Atoi(scanner.Text())
		for j := 0; j < m; j++ {
			scanner.Scan()
			dict[scanner.Text()]++
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	counter := 0
	func() {
		for k, v := range dict {
			if v == n {
				counter++
				defer writer.WriteString(k + "\n")
			}
		}
		writer.WriteString(strconv.Itoa(counter) + "\n")
	}()
	writer.WriteString(strconv.Itoa(len(dict)) + "\n")
	for k, _ := range dict {
		writer.WriteString(k + "\n")
	}
	writer.Flush()
}
