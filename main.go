package main

import (
	"os"

	"github.com/ostracised3rd/advent-of-code-2021/day01"
	"github.com/ostracised3rd/advent-of-code-2021/day02"
	"github.com/ostracised3rd/advent-of-code-2021/day03"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		panic("need day and part to work")
	}

	days(args[0])(args[1])
}

func days(day string) func(part string) {

	dm := map[string]func(part string){
		"day01": day01.Run,
		"day02": day02.Run,
		"day03": day03.Run,
	}

	if fn, ok := dm[day]; ok {
		return fn
	}

	panic("day not found")
}
