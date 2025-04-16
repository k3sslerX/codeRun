package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var xs []int
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp := scanner.Text()
		tmpSlice := strings.Fields(tmp)
		x, _ := strconv.Atoi(tmpSlice[0])
		flag := true
		for _, elem := range xs {
			if elem == x {
				flag = false
				counter++
				break
			}
		}
		if flag {
			xs = append(xs, x)
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(n - counter))
	writer.Flush()
}
