package core

import (
	"testing"
)

func TestInitMoveTables(t *testing.T) {
	InitMoveTables()

	if len(moveTables) != 6 {
		t.Errorf("Initialization of moveTables failed, expected 6 entries but got %d", len(moveTables))
	}
}

func TestNewCube(t *testing.T) {
	c := NewCube()

	for i := 0; i < 8; i++ {
		if c.CornerPermutation[i] != uint8(i) || c.CornerOrientation[i] != 0 {
			t.Errorf("Initialization of corner permutation and orientation failed")
		}
	}

	for i := 0; i < 12; i++ {
		if c.EdgePermutation[i] != uint8(i) || c.EdgeOrientation[i] != 0 {
			t.Errorf("Initialization of edge permutation and orientation failed")
		}
	}
}

func TestSolved(t *testing.T) {
	c := NewCube()

	if !c.solved() {
		t.Error("Newly created cube should be in a solved state")
	}
}

func TestCreateCopyWithmove(t *testing.T) {
	c := NewCube()
	move := Move{
		F: "r",
		N: 1,
	}

	copy := CreateCopyWithmove(c, move)

	if &c == &copy {
		t.Error("CreateCopyWithmove should create a new Cube instance")
	}

	c.Move(move.F, int(move.N))

	for i := 0; i < 8; i++ {
		if c.CornerPermutation[i] != copy.CornerPermutation[i] || c.CornerOrientation[i] != copy.CornerOrientation[i] {
			t.Errorf("Corner permutation and orientation not matching after applying the move")
		}
	}

	for i := 0; i < 12; i++ {
		if c.EdgePermutation[i] != copy.EdgePermutation[i] || c.EdgeOrientation[i] != copy.EdgeOrientation[i] {
			t.Errorf("Edge permutation and orientation not matching after applying the move")
		}
	}
}

func TestMove(t *testing.T) {
	c := NewCube()
	move := Move{
		F: "r",
		N: 1,
	}

	c.Move(move.F, int(move.N))
	expectedCornerPermutation := [8]uint8{0, 1, 2, 3, 4, 5, 7, 6}
	expectedCornerOrientation := [8]uint8{0, 0, 0, 0, 0, 0, 2, 1}

	expectedEdgePermutation := [12]uint8{0, 1, 2, 3, 4, 6, 5, 7, 8, 9, 10, 11}
	expectedEdgeOrientation := [12]uint8{0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0}

	if c.CornerPermutation != expectedCornerPermutation {
		t.Errorf("Corner permutation after 'r' move is not as expected")
	}

	if c.CornerOrientation != expectedCornerOrientation {
		t.Errorf("Corner orientation after 'r' move is not as expected")
	}

	if c.EdgePermutation != expectedEdgePermutation {
		t.Errorf("Edge permutation after r' move is not as expected")
	}

	if c.EdgeOrientation != expectedEdgeOrientation {
		t.Errorf("Edge orientation after 'r' move is not as expected")
	}
}

func TestCreateCopyWithMove(t *testing.T) {
	c := NewCube()
	move := Move{
		F: "l",
		N: 2,
	}

	c2 := CreateCopyWithmove(c, move)

	expectedCornerPermutation := [8]uint8{4, 1, 2, 0, 7, 5, 6, 3}
	expectedCornerOrientation := [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}

	expectedEdgePermutation := [12]uint8{4, 1, 2, 3, 11, 5, 6, 0, 8, 9, 10, 7}
	expectedEdgeOrientation := [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if c2.CornerPermutation != expectedCornerPermutation {
		t.Errorf("Corner permutation after 'l2' move is not as expected")
	}
	if c2.CornerOrientation != expectedCornerOrientation {
		t.Errorf("Corner orientation after 'l2' move is not as expected")
	}

	if c2.EdgePermutation != expectedEdgePermutation {
		t.Errorf("Edge permutation after 'l2' move is not as expected")
	}

	if c2.EdgeOrientation != expectedEdgeOrientation {
		t.Errorf("Edge orientation after 'l2' move is not as expected")
	}
}

func TestPhase1Solved(t *testing.T) {
	c := NewCube()

	if !c.phase1Solved() {
		t.Errorf("Phase 1 should be solved for the initial cube")
	}

	move := Move{
		F: "r",
		N: 1,
	}

	c.MakeMove(move)

	if c.phase1Solved() {
		t.Errorf("Phase 1 should not be solved after 'r' move")
	}
}

func TestCornerOrientationCoordinate(t *testing.T) {
	c := NewCube()
	cornerOrientation := c.CornerOrientationCoordinate()

	if cornerOrientation != 0 {
		t.Errorf("Initial cube's corner orientation coordinate should be 0")
	}

	move := Move{
		F: "r",
		N: 1,
	}
	c.MakeMove(move)
	cornerOrientation = c.CornerOrientationCoordinate()

	if cornerOrientation == 0 {
		t.Errorf("Corner orientation coordinate should not be 0 after 'r' move")
	}
}

func TestEdgeOrientationCoordinate(t *testing.T) {
	c := NewCube()
	edgeOrientation := c.EdgeOrientationCoordinate()

	if edgeOrientation != 0 {
		t.Errorf("Initial cube's edge orientation coordinate should be 0")
	}

	move := Move{
		F: "r",
		N: 1,
	}

	c.MakeMove(move)
	edgeOrientation = c.EdgeOrientationCoordinate()

	if edgeOrientation == 0 {
		t.Errorf("Edge orientation coordinate should not be 0 after 'r' move")
	}
}

func TestUdsCoordinate(t *testing.T) {
	c := NewCube()
	uds := c.UdsCoordinate()

	if uds != 0 {
		t.Errorf("Initial cube's UDS coordinate should be 0")
	}

	move := Move{
		F: "r",
		N: 1,
	}

	c.MakeMove(move)
	uds = c.UdsCoordinate()

	if uds == 0 {
		t.Errorf("UDS coordinate should not be 0 after 'r' move")
	}
}
