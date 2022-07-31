package day02

import "testing"

func TestParseLine(t *testing.T) {
	raw := "forward 5"
	res1, res2 := parseLine(raw)
	want1, want2 := "forward", 5

	if res1 != want1 || res2 != want2 {
		t.Fatalf(`result = %v, %v, want = %v, %v`, res1, res2, want1, want2)
	}
}

func TestMovement(t *testing.T) {
	raw := []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}

	res := movement(raw)
	want := 150

	if res != want {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}
}

func TestMovementWithAim(t *testing.T) {
	raw := []string{"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2"}

	res := movementWithAim(raw)
	want := 900

	if res != want {
		t.Fatalf(`result = %v, want = %v`, res, want)
	}
}
