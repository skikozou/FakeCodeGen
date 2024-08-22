package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		orignFile    string
		fakeCodeFile string
	)
	if len(os.Args) < 3 {
		fmt.Printf("orign file path\n>")
		var text string
		fmt.Scan(&text)
		orignFile = text
		fmt.Printf("fake code file path\n>")
		fmt.Scan(&text)
		fakeCodeFile = text
	} else {
		orignFile = os.Args[1]
		fakeCodeFile = os.Args[2]
	}
	raw, _ := os.ReadFile(orignFile)
	lines := strings.Split(string(raw), "\r\n")
	rawcode := make([][]string, 0)
	for _, l := range lines {
		if l != "" {
			line := strings.Split(l, " ")
			rawcode = append(rawcode, line)
		}
	}
	code := srcGen(rawcode)
	source := code.Build()
	fmt.Println("--------------------------------------------------------------------------------\n" + source + "--------------------------------------------------------------------------------")
	f, _ := os.Create(fakeCodeFile)
	f.Write([]byte(source))
	f.Close()
}
