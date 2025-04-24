package main

import (
	"bufio"
	"os"
	"strconv"
)

func generatePartitions(n int) [][]int {
	var result [][]int
	var current []int
	backtrack(n, n, current, &result)
	return result
}

func backtrack(remaining, maxNum int, current []int, result *[][]int) {
	if remaining == 0 {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := minInt(maxNum, remaining); i >= 1; i-- {
		backtrack(remaining-i, i, append(current, i), result)
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	partitions := generatePartitions(n)

	writer := bufio.NewWriter(os.Stdout)
	for i := len(partitions) - 1; i >= 0; i-- {
		output := ""
		for i, num := range partitions[i] {
			if i > 0 {
				output += " + "
			}
			output += strconv.Itoa(num)
		}
		_, _ = writer.WriteString(output + "\n")
	}
	_ = writer.Flush()
}
