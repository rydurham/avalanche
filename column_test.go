package main

import (
	"fmt"
	"testing"
)

func TestColumnFindsBlockedIndexes(t *testing.T) {
	column := Column{[]string{".", " ", ":", "T"}}

	// We should get true for an index that is out of bounds
	if !column.isIndexBlocked(4) {
		t.Error("Column did not determine correct blockage index")
	}

	// We should get true for an index that is out of bounds
	if !column.isIndexBlocked(10) {
		t.Error("Column did not determine correct blockage index")
	}

	// We should get true for an index that is out of bounds
	if !column.isIndexBlocked(-2) {
		t.Error("Column did not determine correct blockage index")
	}

	// Index 0: "." is not blocked
	if column.isIndexBlocked(0) {
		t.Error("Column flagged an index that is not blocked")
	}

	// Index 1: " " is blocked
	if !column.isIndexBlocked(1) {
		t.Error("Column flagged an index that is blocked")
	}

	// Index 2: ":" is blocked
	if !column.isIndexBlocked(2) {
		t.Error("Column flagged an index that is blocked")
	}

	// Index 3: "T" is blocked
	if !column.isIndexBlocked(3) {
		t.Error("Column did not determine correct blockage index")
	}

	column = Column{[]string{".", ".", "T", " ", " "}}

	if !column.isIndexBlocked(1) {
		t.Error("Column did not use Table as blockage type")
	}
}

func TestColumnCanMoveRocksIntoBlankSpaces(t *testing.T) {
	start := Column{[]string{".", " "}}
	expected := Column{[]string{" ", "."}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("Moving into empty cell failed")
	}
}

func TestColumnWillCollapseRocksIfThereIsNoMoreRoom(t *testing.T) {
	start := Column{[]string{" ", " ", ".", ".", "."}}
	expected := Column{[]string{" ", " ", ".", " ", ":"}}

	start.move(3, 4)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("Collapsing rocks failed")
	}
}

func TestColumnWillNotCollapseRocksIfThereAdjacentSpace(t *testing.T) {
	start := Column{[]string{".", ".", " "}}
	expected := Column{[]string{".", ".", " "}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("Collapsing rocks failed")
	}
}

func TestColumnNotMoveRocksIfAdjacentToBlockedRockStack(t *testing.T) {
	start := Column{[]string{".", ":"}}
	expected := Column{[]string{".", ":"}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("Collapsing rocks failed")
	}
}

func TestColumnWillCollapseRocksIfTheyAreNextToATable(t *testing.T) {
	start := Column{[]string{".", ".", "T"}}
	expected := Column{[]string{" ", ":", "T"}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("Collapsing rocks failed")
	}
}

func TestColumnWillNotMoveATableInTheTargetIndex(t *testing.T) {
	start := Column{[]string{".", "T", " "}}
	expected := Column{[]string{".", "T", " "}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("A target table was moved")
	}
}

func TestColumnWillNotMoveATableInTheSourceIndex(t *testing.T) {
	start := Column{[]string{"T", " ", " "}}
	expected := Column{[]string{"T", " ", " "}}

	start.move(0, 1)

	if !slicesAreEqual(start.cells, expected.cells) {

		fmt.Println(start)
		t.Error("A target table was moved")
	}
}
