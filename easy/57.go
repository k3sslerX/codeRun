package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	counter := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	gen1 := scanner.Text()
	scanner.Scan()
	gen2 := scanner.Text()
	for i := 0; i < len(gen1)-1; i++ {
		found := false
		g1 := string(gen1[i]) + string(gen1[i+1])
		for j := 0; j < len(gen2)-1 && !found; j++ {
			g2 := string(gen2[j]) + string(gen2[j+1])
			if g1 == g2 {
				counter++
				found = true
			}
		}
	}
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(strconv.Itoa(counter))
	writer.Flush()
}
