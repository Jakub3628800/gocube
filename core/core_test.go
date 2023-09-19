package core

import (
	"testing"
)

// Helper function to assert two cubes are equal
func AssertCubesEqual(c Cube, expectedCube Cube, t *testing.T) {
	if c.CornerPermutation != expectedCube.CornerPermutation {
		t.Errorf("CornerPermutation wrong after 'r' move. Expected: %v ; Actual: %v", expectedCube.CornerPermutation, c.CornerPermutation)
	}

	if c.CornerOrientation != expectedCube.CornerOrientation {
		t.Errorf("CornerOrientation wrong after 'r' move. Expected: %v ; Actual: %v", expectedCube.CornerOrientation, c.CornerOrientation)
	}

	if c.EdgePermutation != expectedCube.EdgePermutation {
		t.Errorf("Edgepermutation wrong after 'r' move. Expected: %v ; Actual: %v", expectedCube.EdgePermutation, c.EdgePermutation)
	}

	if c.EdgeOrientation != expectedCube.EdgeOrientation {
		t.Errorf("EdgeOrientation wrong after 'r' move. Expected: %v ; Actual: %v", expectedCube.EdgeOrientation, c.EdgeOrientation)
	}
}

// Test cube constructor
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

func TestMoveR(t *testing.T) {
	c := NewCube()
	InitMoveTables()

	c.Move("r", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{0, 3, 2, 7, 4, 1, 6, 5},
		CornerOrientation: [8]uint8{0, 2, 0, 1, 0, 1, 0, 2},
		EdgePermutation:   [12]uint8{0, 1, 9, 3, 4, 5, 10, 7, 8, 6, 2, 11},
		EdgeOrientation:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	AssertCubesEqual(c, expectedCube, t)

}

func TestMoveL(t *testing.T) {
	c := NewCube()
	InitMoveTables()

	c.Move("l", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{4, 1, 0, 3, 6, 5, 2, 7},
		CornerOrientation: [8]uint8{2, 0, 1, 0, 1, 0, 2, 0},
		EdgePermutation:   [12]uint8{11, 1, 2, 3, 8, 5, 6, 7, 0, 9, 10, 4},
		EdgeOrientation:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	AssertCubesEqual(c, expectedCube, t)

}

func TestMoveU(t *testing.T) {
	c := NewCube()

	InitMoveTables()
	c.Move("u", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{1, 5, 2, 3, 0, 4, 6, 7},
		CornerOrientation: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0},
		EdgePermutation:   [12]uint8{1, 2, 3, 0, 4, 5, 6, 7, 8, 9, 10, 11},
		EdgeOrientation:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	AssertCubesEqual(c, expectedCube, t)
}

func TestMoveD(t *testing.T) {
	c := NewCube()

	InitMoveTables()
	c.Move("d", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{0, 1, 6, 2, 4, 5, 7, 3},
		CornerOrientation: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0},
		EdgePermutation:   [12]uint8{0, 1, 2, 3, 7, 4, 5, 6, 8, 9, 10, 11},
		EdgeOrientation:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	AssertCubesEqual(c, expectedCube, t)
}

func TestMoveF(t *testing.T) {
	c := NewCube()

	InitMoveTables()
	c.Move("f", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{2, 0, 3, 1, 4, 5, 6, 7},
		CornerOrientation: [8]uint8{1, 2, 2, 1, 0, 0, 0, 0},
		EdgePermutation:   [12]uint8{0, 8, 2, 3, 4, 9, 6, 7, 5, 1, 10, 11},
		EdgeOrientation:   [12]uint8{0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0},
	}
	AssertCubesEqual(c, expectedCube, t)
}

func TestMoveB(t *testing.T) {
	c := NewCube()

	InitMoveTables()
	c.Move("b", 1)
	expectedCube := Cube{
		CornerPermutation: [8]uint8{0, 1, 2, 3, 5, 7, 4, 6},
		CornerOrientation: [8]uint8{0, 0, 0, 0, 1, 2, 2, 1},
		EdgePermutation:   [12]uint8{0, 1, 2, 10, 4, 5, 6, 11, 8, 9, 7, 3},
		EdgeOrientation:   [12]uint8{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
	}
	AssertCubesEqual(c, expectedCube, t)
}

type combTestCase struct {
	n        int
	k        int
	expected int
}

func TestComb(t *testing.T) {
	testCases := []combTestCase{
		combTestCase{10, 4, 210},
		combTestCase{5, 3, 10},
		combTestCase{10, 6, 210},
		combTestCase{11, 2, 55},
		combTestCase{77, 4, 1353275},
		combTestCase{10, 0, 1},
		combTestCase{0, 11, 0},
	}

	for _, testCase := range testCases {
		actual := Comb(testCase.n, testCase.k)
		if actual != testCase.expected {
			t.Errorf("Expected %v ; Actual %v", testCase.expected, actual)
		}
	}
}

type coordinateTestCase struct {
	move     string
	expected int
}

func TestCornerOrientationCoordinate(t *testing.T) {
	testCases := []coordinateTestCase{
		coordinateTestCase{"r", 4650},
		coordinateTestCase{"l", 1550},
		coordinateTestCase{"u", 0},
		coordinateTestCase{"d", 0},
		coordinateTestCase{"f", 52},
		coordinateTestCase{"b", 4212},
	}
	for _, testCase := range testCases {
		c := NewCube()
		c.Move(testCase.move, 1)
		actual := CornerOrientationCoordinate(&c)
		if actual != testCase.expected {
			t.Errorf("Expected %v ; Actual %v", testCase.expected, actual)
		}
	}
}

func TestEdgeOrientationCoordinate(t *testing.T) {
	testCases := []coordinateTestCase{
		coordinateTestCase{"r", 0},
		coordinateTestCase{"l", 0},
		coordinateTestCase{"u", 0},
		coordinateTestCase{"d", 0},
		coordinateTestCase{"f", 802},
		coordinateTestCase{"b", 3208},
	}
	for _, testCase := range testCases {
		c := NewCube()
		c.Move(testCase.move, 1)
		actual := EdgeOrientationCoordinate(&c)
		if actual != testCase.expected {
			t.Errorf("Expected %v ; Actual %v", testCase.expected, actual)
		}
	}
}

func TestUdsCoordinate(t *testing.T) {
	testCases := []coordinateTestCase{
		coordinateTestCase{"r", 88},
		coordinateTestCase{"l", 191},
		coordinateTestCase{"u", 0},
		coordinateTestCase{"d", 0},
		coordinateTestCase{"f", 30},
		coordinateTestCase{"b", 285},
	}
	for _, testCase := range testCases {
		c := NewCube()
		c.Move(testCase.move, 1)
		actual := UdsCoordinate(&c)
		if actual != testCase.expected {
			t.Errorf("Expected %v ; Actual %v", testCase.expected, actual)
		}
	}
}

type phase1SolvedTestCase struct {
	move     string
	expected bool
}

func TestPhase1Solved(t *testing.T) {
	testCases := []phase1SolvedTestCase{
		phase1SolvedTestCase{"r", true},
		phase1SolvedTestCase{"l", true},
		phase1SolvedTestCase{"f", true},
		phase1SolvedTestCase{"b", true},
		phase1SolvedTestCase{"u", false},
		phase1SolvedTestCase{"d", false},
	}
	for _, testCase := range testCases {
		c := NewCube()
		c.Move(testCase.move, 1)
		actual := phase1Solved(&c)
		if actual == testCase.expected {
			t.Errorf("phase1solved for %v - expected  %v, got %v", testCase.move, testCase.expected, actual)
		}
	}
}
