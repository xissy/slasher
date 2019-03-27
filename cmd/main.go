package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/xissy/slasher"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	result := slasher.Slasher(line)
	fmt.Println(result)
}
