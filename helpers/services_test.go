package helpers

import "testing"

func TestReverse(t *testing.T) {
	raw := []int{1, 0, 1, 0, 1, 0}
	want := []int{0, 1, 0, 1, 0, 1}
	res := Reverse(raw)

	for i := range want {
		if res[i] != want[i] {
			t.Fatalf(`result = %v, want = %v`, res, want)
		}
	}
}
