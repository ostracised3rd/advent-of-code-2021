package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ostracised3rd/advent-of-code-2021/helpers"
)

func dataParser(raw string) []int {
	dataString := strings.Split(raw, "\n")
	data := []int{}

	for _, i := range dataString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}

		data = append(data, j)
	}

	return data
}

func increaseCounter(inputs []int) int {
	count := 0
	last := inputs[0]
	for _, v := range inputs {
		if v > last {
			count++
		}
		last = v
	}

	return count
}

func threeSums(inputs []int) []int {
	var sums []int

	for i, x := range inputs {
		if i+2 >= len(inputs) {
			break
		}
		sums = append(sums, x+inputs[i+1]+inputs[i+2])
	}

	return sums
}

func p1(raw string) {
	data := dataParser(raw)
	res := increaseCounter(data)
	fmt.Println(res)
}

func p2(raw string) {
	data := dataParser(raw)
	sums := threeSums(data)
	res := increaseCounter(sums)
	fmt.Println(res)
}

func Run(part string) {
	raw := helpers.ReadData("assets/d01.txt")

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
