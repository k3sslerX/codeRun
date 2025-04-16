package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	nums := make([]int, n)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(parts[i])
	}
	cost := 0.0
	for len(nums) > 1 {
		minInd1 := 0
		minInd2 := 1
		for i := 0; i < len(nums); i++ {
			if i == minInd1 {
				continue
			}
			if nums[i] < nums[minInd1] {
				minInd2 = minInd1
				minInd1 = i
			} else if nums[i] < nums[minInd2] {
				minInd2 = i
			}
		}
		sum := nums[minInd1] + nums[minInd2]
		cost += 0.05 * float64(sum)
		nums[minInd1] = sum
		nums[0], nums[minInd2] = nums[minInd2], nums[0]
		nums = nums[1:]
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.FormatFloat(cost, 'f', 2, 64))
	writer.Flush()
}
