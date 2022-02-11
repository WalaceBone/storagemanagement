package main

import (
	"errors"
	"strconv"
	"strings"
)

type Parser struct {
	warehouse *Warehouse
}

func parseInt(str string) (uint, error) {
	for index := 0; index < len(str); index += 1 {
		if str[index] < '0' || str[index] > '9' {
			return 0, errors.New("Invalid integer")
		}
	}
	intValue, err := strconv.Atoi(str)
	return uint(intValue), err
}

func (p Parser) parseSettings(lines []string) error {
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
			p.warehouse.Size = Size{xPos, yPos}
			p.warehouse.Lifetime = lifetime
		}
		switch len(params) {
		case 3:
			p.warehouse.Forklifts = append(p.warehouse.Forklifts, Forklift{pos, Package{}, FStatus("WAIT"), params[0]})
		case 4:
			weight, err := parseInt(params[3])
			if err != nil {
				return err
			}
			p.warehouse.Packages = append(p.warehouse.Packages, NewPackage(weight, xPos, yPos, params[0]))
		case 5:
			capacity, err := parseInt(params[3])
			if err != nil {
				return err
			}
			cooldown, err := parseInt(params[4])
			p.warehouse.Trucks = append(p.warehouse.Trucks, NewTruck(cooldown, xPos, yPos, capacity, params[0]))
		}
	}
	return nil
}
