package main

import "testing"

func TestSliceContains(t *testing.T) {
	slice := []string{"a", "e", "i", "o", "u"}

	if !contains(slice, "o") {
		t.Error("Could not reliably determine slice containment status")
	}

	if contains(slice, "x") {
		t.Error("Could not reliably determine slice containment status")
	}
}

func TestFindInSlice(t *testing.T) {
	slice := []string{"a", "e", "i", "o", "u", "u"}

	if find(slice, "o") != 3 {
		t.Error("Did not find expected content in slice")
	}

	if find(slice, "x") != 0 {
		t.Error("Found unexpected conteint in slice")
	}

	if find(slice, "u") != 4 {
		t.Error("Unexpected index returned from find operation")
	}
}

func TestCheckSliceEquality(t *testing.T) {
	sliceA := []string{"A", "B", "C"}
	sliceB := []string{"X", "Y", "Z"}
	sliceC := []string{"A", "B", "C", "D"}
	sliceD := []string{"A", "B", "C"}

	if slicesAreEqual(sliceA, sliceB) {
		t.Error("Slice comparison failed")
	}

	if slicesAreEqual(sliceA, sliceC) {
		t.Error("Slice comparison failed")
	}

	if !slicesAreEqual(sliceA, sliceD) {
		t.Error("Slice comparison failed")
	}
}
