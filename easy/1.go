package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func sort(slice *[]int) {
	for i := 0; i < len(*slice); i++ {
		for j := i; j < len(*slice); j++ {
			if (*slice)[j] < (*slice)[i] {
				(*slice)[i], (*slice)[j] = (*slice)[j], (*slice)[i]
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	nums := make([]int, 3)
	nums[0], _ = strconv.Atoi(parts[0])
	nums[1], _ = strconv.Atoi(parts[1])
	nums[2], _ = strconv.Atoi(parts[2])
	writer := bufio.NewWriter(os.Stdout)
	sort(&nums)
	writer.WriteString(strconv.Itoa(nums[1]))
	writer.Flush()
}
