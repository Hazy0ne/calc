package main

import (
	"bufio"
	"calc/app"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hi, go count")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", " ")
		app.Base(strings.ToUpper(strings.TrimSpace(s)))
	}
}
