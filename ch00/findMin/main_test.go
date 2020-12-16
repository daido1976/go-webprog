package main

import (
	"testing"
)

func TestFindMin(t *testing.T) {
	expect := 3
	actual := findMin([]int{10, 3, 5, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}
