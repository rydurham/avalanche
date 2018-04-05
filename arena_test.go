package main

import "testing"

func TestArenaToStringMultiColumn(t *testing.T) {
	arena := Arena{[]Column{
		Column{[]string{".", " ", ":", "T"}},
		Column{[]string{".", " ", ":", "T"}},
	}}

	expected := "+--+\n|..|\n|  |\n|::|\n|TT|\n+--+\n"

	if arena.toString() != expected {
		t.Error(arena.toString())
	}
}

func TestArenaToStringSingleColumn(t *testing.T) {
	arena := Arena{[]Column{
		Column{[]string{".", " ", ":", "T"}},
	}}

	expected := ".\n \n:\nT\n"

	if arena.toString() != expected {
		t.Error("Arena toString() returned unexpected output")
	}
}

func TestArenaDuplicationCreatesDeepCopy(t *testing.T) {
	arena := Arena{[]Column{
		Column{[]string{" ", ".", ":", "T"}},
	}}

	duplicate := arena.duplicate()
	arena.columns[0].cells[0] = "X"

	if arena.columns[0].cells[0] == duplicate.columns[0].cells[0] {
		t.Error("Arena duplication failed: deep copy not performed")
	}
}

func TestArenaDuplicationRowsMatch(t *testing.T) {
	arena := Arena{[]Column{
		Column{[]string{" ", ".", ":", "T"}},
	}}

	duplicate := arena.duplicate()

	if arena.numberOfRows() != duplicate.numberOfRows() {
		t.Error("Arena duplication failed: row counts do not match")
	}
}

func TestArenaDuplicationColumnsMatch(t *testing.T) {
	arena := Arena{[]Column{
		Column{[]string{" ", ".", ":", "T"}},
	}}

	duplicate := arena.duplicate()

	if arena.numberOfColumns() != duplicate.numberOfColumns() {
		t.Error("Arena duplication failed: column counts do not match")
	}
}

func TestGetArenaColumnCount(t *testing.T) {
	arena0 := Arena{[]Column{}}
	arena1 := Arena{[]Column{
		Column{[]string{" ", ".", ":", "T"}},
	}}

	if arena0.numberOfColumns() != 0 {
		t.Error("Arena0 column count failed")
	}

	if arena1.numberOfColumns() != 1 {
		t.Error("Arena1 column count failed")
	}
}

func TestGetArenaRowCount(t *testing.T) {
	arena0 := Arena{[]Column{}}
	arena1 := Arena{[]Column{
		Column{[]string{" ", ".", ":", "T"}},
	}}

	if arena0.numberOfRows() != 0 {
		t.Error("Arena0 row count failed")
	}

	if arena1.numberOfRows() != 4 {
		t.Error("Arena1 row count failed")
	}
}
