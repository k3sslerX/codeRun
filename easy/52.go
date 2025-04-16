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
	var key, value string
	dict := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		tmp := strings.Fields(scanner.Text())
		key = tmp[0]
		value = tmp[1]
		dict[key] = value
	}
	scanner.Scan()
	search := scanner.Text()
	writer := bufio.NewWriter(os.Stdout)
	for k, v := range dict {
		if k == search {
			_, _ = writer.WriteString(v)
			break
		} else if v == search {
			_, _ = writer.WriteString(k)
			break
		}
	}
	_ = writer.Flush()
}
