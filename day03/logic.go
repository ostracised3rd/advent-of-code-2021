package day03

import (
	"fmt"
	"math"
	"strings"

	"github.com/ostracised3rd/advent-of-code-2021/helpers"
)

func dataParser(raw string) [][]int {
	var data [][]int

	for _, line := range strings.Split(raw, "\n") {
		var datum []int
		for _, char := range []rune(line) {
			if char == '1' {
				datum = append(datum, 1)
			} else {
				datum = append(datum, 0)
			}
		}

		data = append(data, datum)
	}

	return data
}

func conversion(data []int) int {
	acc := 0
	for i, x := range helpers.Reverse(data) {
		acc += int(math.Pow(float64(2), float64(i))) * x
	}

	return acc
}

func population(data string) (int, []int) {
	i := 0
	acc := []int{}

	for _, x := range dataParser(data) {
		i += 1
		if len(acc) == 0 {
			acc = x
		} else {
			for j, y := range x {
				acc[j] += y
			}
		}
	}

	return i, acc
}

func indexPop(index int, data [][]int) int {
	i := 0
	for _, x := range data {
		i += x[index]
	}
	return i
}

func filterPop(index int, val int, data [][]int) [][]int {
	var res [][]int

	for _, x := range data {

		if x[index] == val {
			res = append(res, x)
		}
	}

	return res
}

func chemicalValue(val [2]int, index int, data [][]int) []int {
	if len(data) == 1 {
		return data[0]
	}

	mid := int(math.Ceil(float64(len(data)) / 2.0))
	pop := indexPop(index, data)
	var keep int

	if pop >= mid {
		keep = val[0]
	} else {
		keep = val[1]
	}

	return chemicalValue(val, index+1, filterPop(index, keep, data))
}

func lifeSupport(raw string) int {
	data := dataParser(raw)
	oxy := chemicalValue([2]int{1, 0}, 0, data)
	co2 := chemicalValue([2]int{0, 1}, 0, data)

	return conversion(oxy) * conversion(co2)
}

func powerConsumption(data string) int {
	len, pop := population(data)

	var gamma []int
	var epsilon []int
	mid := len / 2

	for _, x := range pop {
		if x > mid {
			gamma = append(gamma, 1)
			epsilon = append(epsilon, 0)
		} else {
			gamma = append(gamma, 0)
			epsilon = append(epsilon, 1)
		}
	}

	return conversion(gamma) * conversion(epsilon)
}

func P1() {
	raw := helpers.ReadData("assets/d03.txt")
	res := powerConsumption(raw)
	fmt.Println(res)
}

func P2() {
	raw := helpers.ReadData("assets/d03.txt")
	res := lifeSupport(raw)
	fmt.Println(res)
}
