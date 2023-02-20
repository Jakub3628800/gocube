package core

import (
	"cube/utils"
	"fmt"
)

var moveTables map[string][4][4]uint8

func initMoveTables() {
	moveTables = make(map[string][4][4]uint8)

	moveTables["r"] = [4][4]uint8{{1, 3, 7, 5}, {9, 6, 10, 2}, {2, 1, 2, 1}, {0, 0, 0, 0}}
	moveTables["l"] = [4][4]uint8{{0, 4, 6, 2}, {8, 0, 11, 4}, {2, 1, 2, 1}, {0, 0, 0, 0}}
	moveTables["u"] = [4][4]uint8{{1, 5, 4, 0}, {1, 2, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	moveTables["d"] = [4][4]uint8{{3, 2, 6, 7}, {5, 4, 7, 6}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	moveTables["f"] = [4][4]uint8{{0, 2, 3, 1}, {1, 8, 5, 9}, {1, 2, 1, 2}, {1, 1, 1, 1}}
	moveTables["b"] = [4][4]uint8{{4, 5, 7, 6}, {3, 10, 7, 11}, {1, 2, 1, 2}, {1, 1, 1, 1}}

}

type cube struct {
	cornerPermutation [8]uint8
	cornerOrientation [8]uint8
	edgePermutation   [12]uint8
	edgeOrientation   [12]uint8
}

func (c *cube) initCube() {
	c.cornerPermutation = [8]uint8{0, 1, 2, 3, 4, 5, 6, 7}
	c.cornerOrientation = [8]uint8{0, 0, 0, 0, 0, 0, 0, 0}
	c.edgePermutation = [12]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	c.edgeOrientation = [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
}

// Return hash of cube object.
func (c *cube) hash() string {
	// hash corner orientation

	return "s"
}

func (c *cube) cornerOrientationCoordinate() int {
	result := 0
	multiplier := 1

	for i := 0; i < 8; i++ {
		result += int(c.cornerOrientation[i]) * multiplier
		multiplier = multiplier * 3
	}
	return result
}

func (c *cube) edgeOrientationCoordinate() int {
	result := 0
	multiplier := 1
	for i := 0; i < 12; i++ {
		result += int(c.edgeOrientation[i]) * multiplier
		multiplier = multiplier * 2
	}
	return result
}

func (c *cube) udsCoordinate() int {
	k := -1
	result := 0
	var val uint8
	for i := 0; i < len(c.edgePermutation); i++ {
		val = c.edgePermutation[i]
		if val > 7 {
			k += 1

		} else if k > 0 {
			result += utils.Comb(i, k)
		}

	}
	return result
}

// Is the cube solved?
func (c *cube) solved() bool {
	return true
}

// Apply move to cube
func (cb *cube) move(move string, nr int) {

	var t uint8
	c := moveTables[move][0]
	e := moveTables[move][1]
	ct := moveTables[move][2]
	et := moveTables[move][3]

	for i := 0; i < nr; i++ {

		t = cb.cornerPermutation[c[0]]
		cb.cornerPermutation[c[0]] = cb.cornerPermutation[c[1]]
		cb.cornerPermutation[c[1]] = cb.cornerPermutation[c[2]]
		cb.cornerPermutation[c[2]] = cb.cornerPermutation[c[3]]
		cb.cornerPermutation[c[3]] = t

		t = cb.edgePermutation[e[0]]
		cb.edgePermutation[e[0]] = cb.edgePermutation[e[1]]
		cb.edgePermutation[e[1]] = cb.edgePermutation[e[2]]
		cb.edgePermutation[e[2]] = cb.edgePermutation[e[3]]
		cb.edgePermutation[e[3]] = t

		t = cb.cornerOrientation[c[0]]
		cb.cornerOrientation[c[0]] = ((cb.cornerOrientation[c[1]] + ct[0]) % 3)
		cb.cornerOrientation[c[1]] = ((cb.cornerOrientation[c[2]] + ct[1]) % 3)
		cb.cornerOrientation[c[2]] = ((cb.cornerOrientation[c[3]] + ct[2]) % 3)
		cb.cornerOrientation[c[3]] = ((t + ct[3]) % 3)

		t = cb.edgeOrientation[e[0]]
		cb.edgeOrientation[e[0]] = (cb.edgeOrientation[e[1]] + et[0]) % 2
		cb.edgeOrientation[e[1]] = (cb.edgeOrientation[e[2]] + et[1]) % 2
		cb.edgeOrientation[e[2]] = (cb.edgeOrientation[e[3]] + et[2]) % 2
		cb.edgeOrientation[e[3]] = (t + et[3]) % 2
	}
}

func (c *cube) printcube() {
	fmt.Println("================================")
	fmt.Println("cper", c.cornerPermutation)
	fmt.Println("cor ", c.cornerOrientation)
	fmt.Println("eper", c.edgePermutation)
	fmt.Println("eor ", c.edgeOrientation)
	fmt.Println()
}
