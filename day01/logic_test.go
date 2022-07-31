package day01

import "testing"

func TestDataParser(t *testing.T) {
	raw := "1721\n979\n366\n299\n675\n1456"
	want := []int{1721, 979, 366, 299, 675, 1456}

	data := dataParser(raw)

	for i := range want {
		if want[i] != data[i] {
			t.Fatalf(`dataParser = %v, want = %v`, data, want)
		}
	}
}

func TestIncreaseCounter(t *testing.T) {
	raw := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := 7

	res := increaseCounter(raw)

	if res != want {
		t.Fatalf(`result was: %d, wanted: %d`, res, want)
	}
}

func TestThreeSums(t *testing.T) {
	raw := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	want := []int{607, 618, 618, 617, 647, 716, 769, 792}
	res := threeSums(raw)

	if len(want) != len(res) {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}

	for i := range want {
		if want[i] != res[i] {
			t.Fatalf(`result = %v, want = %v`, res, want)
		}
	}
}
