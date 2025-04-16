package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var pairs []string
	counterLie := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		pair := scanner.Text()
		flag := true
		for _, elem := range pairs {
			if elem == pair {
				flag = false
				counterLie++
				break
			}
		}
		pairsSlice := strings.Fields(pair)
		a, _ := strconv.Atoi(pairsSlice[0])
		b, _ := strconv.Atoi(pairsSlice[1])
		if a < 0 || b < 0 {
			flag = false
			counterLie++
		} else if a+b+1 != n {
			flag = false
			counterLie++
		}
		if flag {
			pairs = append(pairs, pair)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(n - counterLie))
	writer.Flush()
}
