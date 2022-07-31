package day03

import (
	"testing"
)

func TestDataParser(t *testing.T) {
	raw := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
	res := dataParser(raw)
	want := [][]int{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}

	if len(want) != len(res) {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}

	for i, sub := range want {
		for j := range sub {
			if sub[j] != res[i][j] {
				t.Fatalf(`result = %v, want = %v`, res, want)
			}
		}
	}
}

func TestPopulation(t *testing.T) {
	raw := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
	res1, res2 := population(raw)
	want1 := 12
	want2 := []int{7, 5, 8, 7, 5}

	if res1 != want1 {
		t.Fatalf(`result = %v, want = %v`, res1, want1)
	}

	for i := range want2 {
		if want2[i] != res2[i] {
			t.Fatalf(`result = %v, want = %v`, res2, want2)
		}
	}
}

func TestPowerConsumption(t *testing.T) {
	raw := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
	want := 198
	res := powerConsumption(raw)

	if want != res {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}
}

func TestChemicalValue(t *testing.T) {
	raw := [][]int{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}
	want := []int{1, 0, 1, 1, 1}
	res := chemicalValue([2]int{1, 0}, 0, raw)

	for i := range want {
		if res[i] != want[i] {
			t.Fatalf(`result = %v, want = %v`, res, want)
		}
	}
}

func TestLifeSupport(t *testing.T) {
	raw := "00100\n11110\n10110\n10111\n10101\n01111\n00111\n11100\n10000\n11001\n00010\n01010"
	want := 230
	res := lifeSupport(raw)

	if res != want {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}
}

func TestConversion(t *testing.T) {
	raw := []int{1, 0, 1, 0, 1, 0}
	res := conversion(raw)
	want := 42

	if res != want {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}
}
