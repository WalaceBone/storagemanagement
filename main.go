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
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()

	parser := Parser{}
	warehouse, err := parser.parseSettings(lines)
	if err != nil {
		fmt.Errorf("error : %s\n", err)
	}
	//warehouse.Dump()
	Simulation(warehouse)
}
