package main

import (
	"bufio"
	"fmt"
	"os"
)

/*func main() {

	i := [3]int{4, 2, 1}
	var p *int

	p = &i[0]
	fmt.Printf("%d\n", i[0])
	fmt.Printf("%d\n", *p)
	i[0] = 2
	fmt.Printf("%d\n", i[0])
	fmt.Printf("%d\n", *p)
	*p = 7

	fmt.Printf("%d\n", i[0])
	fmt.Printf("%d\n", *p)

	packages := [1]Package{}
	packages[0] = NewPackage(0, 0, 0, "Toto")

	var w Cell

	w.P = &packages[0]

	packages[0].Dump()
	w.P.Dump()

	packages[0].Name = "LULU"

	packages[0].Dump()
	w.P.Dump()
}*/

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
	/*	warehouse.Dump()
		warehouse.Forklifts[0].Dump()
		warehouse.move(2, &warehouse.Forklifts[0])
		warehouse.Forklifts[0].Dump()
		warehouse.DumpMap()*/

	warehouse.Simulation()
}
