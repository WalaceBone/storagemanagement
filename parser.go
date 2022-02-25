package main

import (
	"errors"
	"strconv"
	"strings"
)

type Parser struct {
}

func parseWeight(w string) int {
	w = strings.ToUpper(w)
	if w == "GREEN" {
		return GREEN
	}
	if w == "YELLOW" {
		return YELLOW
	}
	if w == "BLUE" {
		return BLUE
	}
	return -1
}

func parseInt(str string) (int, error) {
	for index := 0; index < len(str); index += 1 {
		if str[index] < '0' || str[index] > '9' {
			return 0, errors.New("Invalid integer")
		}
	}
	intValue, err := strconv.Atoi(str)
	return int(intValue), err
}

func (p Parser) parseSettings(lines []string) (*Warehouse, error) {
	w := Warehouse{}
	for index, line := range lines {
		params := strings.Split(line, " ")
		if index == 0 {
			x, err := parseInt(params[0])
			if err != nil {
				return nil, err
			}
			y, err := parseInt(params[1])
			if err != nil {
				return nil, err
			}
			lifetime, err := parseInt(params[2])
			if err != nil {
				return nil, err
			}
			w = NewWarehouse(x, y, lifetime)
		} else {
			switch len(params) {
			case 3:
				x, err := parseInt(params[1])
				if err != nil {
					return nil, err
				}
				y, err := parseInt(params[2])
				if err != nil {
					return nil, err
				}
				w.GetCellIDFromPosition(x, y)
				err = w.addForklift(NewForklift(w.GetCellIDFromPosition(x, y), x, y, params[0]))
				if err != nil {
					return nil, err
				}
				//p.warehouse.Forklifts = append(p.warehouse.Forklifts, NewForklift(pos.x, pos.y, params[0]))
			case 4:
				weight := parseWeight(params[3])
				if weight == -1 {
					return nil, errors.New("invalid Weight")
				}
				x, err := parseInt(params[1])
				if err != nil {
					return nil, err
				}
				y, err := parseInt(params[2])
				if err != nil {
					return nil, err
				}
				err = w.addPackage(NewPackage(w.GetCellIDFromPosition(x, y), weight, x, y, params[0]))
				if err != nil {
					return nil, err
				}
				//p.warehouse.Packages = append(p.warehouse.Packages, )
			case 5:
				capacity, err := parseInt(params[3])
				if err != nil {
					return nil, err
				}
				cooldown, err := parseInt(params[4])
				x, err := parseInt(params[1])
				if err != nil {
					return nil, err
				}
				y, err := parseInt(params[2])
				if err != nil {
					return nil, err
				}
				err = w.addTruck(NewTruck(w.GetCellIDFromPosition(x, y), cooldown, x, y, capacity, params[0]))
				if err != nil {
					return nil, err
				}
				//p.warehouse.Trucks = append(p.warehouse.Trucks, )
			}
		}
	}
	return &w, nil
}
