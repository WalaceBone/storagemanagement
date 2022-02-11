package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Program takes one argument: path of the file")
		return
	}
	path := args[1]

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	var wh = Warehouse{}
	parser := Parser{&wh}
	parser.parseSettings(lines)
}
