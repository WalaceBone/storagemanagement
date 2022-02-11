package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInt(str string) (int, error) {
	for index := 0; index < len(str); index += 1 {
		if str[index] < '0' || str[index] > '9' {
			return 0, errors.New("Invalid integer")
		}
	}
	return strconv.Atoi(str)
}

func parseSettings(warehouse *Warehouse, lines []string) error {
	for index, line := range lines {
		params := strings.Split(line, " ")
		xPos, err := parseInt(params[0])
		if err != nil {
			return err
		}
		yPos, err := parseInt(params[1])
		if err != nil {
			return err
		}
		pos := Position{xPos, yPos}
		if index == 0 {
			lifetime, err := parseInt(params[2])
			if err != nil {
				return err
			}
			warehouse.Size = Size{xPos, yPos}
			warehouse.Lifetime = uint(lifetime)
		}
		switch len(params) {
		case 3:
			warehouse.Forklifts = append(warehouse.Forklifts, Forklift{pos, Package{}, FStatus("WAIT"), params[0]})
		case 4:
			weightInt, err := parseInt(params[3])
			if err != nil {
				return err
			}
			weight := weight(weightInt)
			warehouse.Packages = append(warehouse.Packages, Package{weight, pos, params[0]})
		case 5:
			//warehouse.Trucks = append(warehouse.Trucks, Truck{pos, Package{}, FStatus("WAIT"), params[0]})
		}
	}
	return nil
}

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
	parseSettings(&wh, lines)
}
