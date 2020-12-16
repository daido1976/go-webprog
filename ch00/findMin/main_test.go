package main

import (
	"reflect"
	"testing"
)

func TestFindMin(t *testing.T) {
	expect := 3
	actual := findMin([]int{10, 3, 5, 7})
	if actual != expect {
		t.Errorf(`expect="%d" actual="%d"`, expect, actual)
	}
}

func TestFindMinWithIdx(t *testing.T) {
	expect := map[string]int{"idx": 1, "value": 3}
	actual := findMinWithIdx([]int{10, 3, 5, 7})
	if !reflect.DeepEqual(expect, actual) {
		t.Error("actual:", actual, "expect:", expect)
	}
}
