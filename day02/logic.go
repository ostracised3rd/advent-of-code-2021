package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ostracised3rd/advent-of-code-2021/helpers"
)

func parseLine(line string) (string, int) {
	temp := strings.Split(line, " ")
	cmd := temp[0]
	val, err := strconv.Atoi(strings.Trim(temp[1], " "))

	if err != nil {
		panic(err.Error())
	}

	return cmd, val
}

func movement(data []string) int {
	x, y := 0, 0

	for _, v := range data {
		cmd, val := parseLine(v)

		switch cmd {
		case "down":
			y += val
		case "forward":
			x += val
		case "up":
			y -= val
		}
	}

	return x * y
}

func movementWithAim(data []string) int {
	x, y, aim := 0, 0, 0

	for _, v := range data {
		cmd, val := parseLine(v)

		switch cmd {
		case "down":
			aim += val
		case "forward":
			x += val
			y += (val * aim)
		case "up":
			aim -= val
		}
	}

	return x * y
}

func p1(raw string) {
	data := strings.Split(raw, "\n")
	res := movement(data)
	fmt.Println(res)
}

func p2(raw string) {
	data := strings.Split(raw, "\n")
	res := movementWithAim(data)
	fmt.Println(res)
}

func Run(part string) {
	raw := helpers.ReadData("assets/d02.txt")

	dm := map[string]func(raw string){
		"p1": p1,
		"p2": p2,
	}

	if fn, ok := dm[part]; ok {
		fn(raw)
	} else {
		panic("part not found")
	}
}
