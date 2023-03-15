package tables

import (
	"cube/core"
	"testing"
)

func TestTable(t *testing.T) {
	initphase1Table()
	createTablePhase1()
}

func TestSaveLoadPhase1Table(t *testing.T) {
	initphase1Table()
	table, err := loadPhase1Table("phase1.gob")
	if err != nil {
		t.Fatalf("Failed to load phase1 table: %v", err)
	}

	if table == nil {
		t.Fatal("Loaded phase1 table is nil")
	}
}

func TestInsertPhase1TableItem(t *testing.T) {
	table := &phase1Table{}
	cube := core.NewCube()
	move := core.Move{"r", 1}
	previousMove := core.Move{"l", 1}
	moveItem := MoveItem{
		cube:             cube,
		lastMove:         move,
		previousLastMove: previousMove,
		movesLeft:        3,
	}

	insertphase1TableItem(moveItem, 1, table)

	cor := cube.CornerOrientationCoordinate()
	eor := cube.EdgeOrientationCoordinate()
	udslice := cube.UdsCoordinate()

	if table[cor][eor][udslice] != 1 {
		t.Errorf("Expected table item value to be 1, but got %d", table[cor][eor][udslice])
	}
}

func TestOppositeMove(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"r", "l"},
		{"f", "b"},
		{"u", "d"},
		{"l", "r"},
		{"b", "f"},
		{"d", "u"},
	}

	for _, tc := range testCases {
		result := oppositeMove(tc.input)
		if result != tc.expected {
			t.Errorf("Expected opposite of %s to be %s, but got %s", tc.input, tc.expected, result)
		}
	}
}
