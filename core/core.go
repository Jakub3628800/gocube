package core

import "fmt"

// Cube state
type Cube struct {
	CornerPermutation [8]uint8
	CornerOrientation [8]uint8
	EdgePermutation   [12]uint8
	EdgeOrientation   [12]uint8
}

func NewCube() Cube {
	return Cube{
		CornerPermutation: [8]uint8{0, 1, 2, 3, 4, 5, 6, 7},
		CornerOrientation: [8]uint8{0, 0, 0, 0, 0, 0, 0, 0},
		EdgePermutation:   [12]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		EdgeOrientation:   [12]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// Move that changes cube state
type Move struct {
	F string // r, l, f, b, u, d
	N uint8  // 1, 2, 3

}

var moveTables map[string][4][4]uint8

// moveTables decribe how the cube changes when a move is applied
func InitMoveTables() {
	moveTables = make(map[string][4][4]uint8)

	moveTables["r"] = [4][4]uint8{{1, 3, 7, 5}, {9, 6, 10, 2}, {2, 1, 2, 1}, {0, 0, 0, 0}}
	moveTables["l"] = [4][4]uint8{{0, 4, 6, 2}, {8, 0, 11, 4}, {2, 1, 2, 1}, {0, 0, 0, 0}}
	moveTables["u"] = [4][4]uint8{{1, 5, 4, 0}, {1, 2, 3, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	moveTables["d"] = [4][4]uint8{{3, 2, 6, 7}, {5, 4, 7, 6}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	moveTables["f"] = [4][4]uint8{{0, 2, 3, 1}, {1, 8, 5, 9}, {1, 2, 1, 2}, {1, 1, 1, 1}}
	moveTables["b"] = [4][4]uint8{{4, 5, 7, 6}, {3, 10, 7, 11}, {1, 2, 1, 2}, {1, 1, 1, 1}}

}

// Apply a move to cube
func (cb *Cube) Move(move string, nr int) {

	var t uint8
	c := moveTables[move][0]
	e := moveTables[move][1]
	ct := moveTables[move][2]
	et := moveTables[move][3]

	for i := 0; i < nr; i++ {
		t = cb.CornerPermutation[c[0]]
		cb.CornerPermutation[c[0]] = cb.CornerPermutation[c[1]]
		cb.CornerPermutation[c[1]] = cb.CornerPermutation[c[2]]
		cb.CornerPermutation[c[2]] = cb.CornerPermutation[c[3]]
		cb.CornerPermutation[c[3]] = t

		t = cb.EdgePermutation[e[0]]
		cb.EdgePermutation[e[0]] = cb.EdgePermutation[e[1]]
		cb.EdgePermutation[e[1]] = cb.EdgePermutation[e[2]]
		cb.EdgePermutation[e[2]] = cb.EdgePermutation[e[3]]
		cb.EdgePermutation[e[3]] = t

		t = cb.CornerOrientation[c[0]]
		cb.CornerOrientation[c[0]] = ((cb.CornerOrientation[c[1]] + ct[0]) % 3)
		cb.CornerOrientation[c[1]] = ((cb.CornerOrientation[c[2]] + ct[1]) % 3)
		cb.CornerOrientation[c[2]] = ((cb.CornerOrientation[c[3]] + ct[2]) % 3)
		cb.CornerOrientation[c[3]] = ((t + ct[3]) % 3)

		t = cb.EdgeOrientation[e[0]]
		cb.EdgeOrientation[e[0]] = (cb.EdgeOrientation[e[1]] + et[0]) % 2
		cb.EdgeOrientation[e[1]] = (cb.EdgeOrientation[e[2]] + et[1]) % 2
		cb.EdgeOrientation[e[2]] = (cb.EdgeOrientation[e[3]] + et[2]) % 2
		cb.EdgeOrientation[e[3]] = (t + et[3]) % 2
	}
}

// Calculate binomial coefficient -  n choose k
func Comb(n int, k int) int {
	if k == 0 {
		return 1
	}
	return (n * Comb(n-1, k-1)) / k
}

func CornerOrientationCoordinate(c *Cube) int {
	result := 0
	multiplier := 1

	for i := 0; i < 8; i++ {
		result += int(c.CornerOrientation[i]) * multiplier
		multiplier = multiplier * 3
	}
	return result
}

func EdgeOrientationCoordinate(c *Cube) int {
	result := 0
	multiplier := 1
	for i := 0; i < 12; i++ {
		result += int(c.EdgeOrientation[i]) * multiplier
		multiplier = multiplier * 2
	}
	return result
}

func UdsCoordinate(c *Cube) int {
	k := -1
	result := 0
	var val uint8
	for i := 0; i < len(c.EdgePermutation); i++ {
		val = c.EdgePermutation[i]
		if val > 7 {
			k += 1

		} else if k > 0 {
			result += Comb(i, k)
		}

	}
	return result
}

func phase1Solved(c *Cube) bool {
	return (CornerOrientationCoordinate(c) == 0) && (EdgeOrientationCoordinate(c) == 0)
}

func PrintCube(c *Cube) {
	fmt.Println("=== Cube State =================")
	fmt.Println("cper", c.CornerPermutation)
	fmt.Println("cor ", c.CornerOrientation)
	fmt.Println("eper", c.EdgePermutation)
	fmt.Println("eor ", c.EdgeOrientation)
	fmt.Println("================================")
	fmt.Println()
}

func PrintCoordinates(c *Cube) {
	fmt.Println("=== Coordinates ================")
	fmt.Println("cor coordinate", CornerOrientationCoordinate(c))
	fmt.Println("eor coordinate", EdgeOrientationCoordinate(c))
	fmt.Println("uds coordinate", UdsCoordinate(c))
	fmt.Println("phase1 solved", phase1Solved(c))
	fmt.Println("================================")
	fmt.Println()
}

//func (cb *Cube) MakeMove(move Move) Cube {
//	var result Cube
//	result.Move(move.F, int(move.N))
//	return result
//}

//// return hash of cube object.
//func (c *Cube) hash() string {
//	// hash corner orientation
//
//	return "s"
//}

// // is the cube solved?  TODO: not implemented
// func (c *Cube) solved() bool {
// 	return true
// }

//func createCopy(src Cube) Cube {
//	dst := Cube{
//		CornerPermutation: src.CornerPermutation,
//		CornerOrientation: src.CornerOrientation,
//		EdgePermutation:   src.EdgePermutation,
//		EdgeOrientation:   src.EdgeOrientation,
//	}
//	return dst
//}
//
//func CreateCopyWithmove(src Cube, move Move) Cube {
//	copy := createCopy(src)
//	copy.MakeMove(move)
//	return copy
//}

//func GenerateTablesPhase1() {
//
//	var phase1table [2186][2047][494]uint8
//}

//var p1Moves []Move
//
//func InitP1Moves() {
//	p1Moves = []Move{
//		Move{"r", 2},
//		Move{"l", 2},
//		Move{"f", 2},
//		Move{"b", 2},
//		Move{"u", 1},
//		Move{"u", 2},
//		Move{"u", 3},
//		Move{"d", 1},
//		Move{"d", 2},
//		Move{"d", 3},
//	}
//}

//func SolvePhase1(c *Cube, moves []Move, movesLeft int) ([]Move, bool) {
//	//fmt.Println(moves, movesLeft)
//	if movesLeft == 0 {
//		if c.phase1Solved() {
//			return moves, true
//		}
//		return moves, false
//	}
//
//	var solved bool
//	results := [][]Move{}
//
//	for i := 0; i < len(p1Moves); i++ {
//		fmt.Println(p1Moves[i], movesLeft, moves)
//		moves, solved = SolvePhase1(CreateCopyWithmove(c, p1Moves[i]), append(moves, p1Moves[i]), movesLeft-1)
//		if solved {
//			results = append(results, append(moves, p1Moves[i]))
//		}
//	}
//
//	if len(results) > 0 {
//		return results[0], true
//	}
//	return nil, false
//
//}
//
//func BreadthPhase1(c *Cube, moves []Move, movesLeft int) {
//	var solved bool
//
//	results := [][]Move{}
//	for i := 0; i < len(p1Moves); i++ {
//		moves, solved = SolvePhase1(CreateCopyWithmove(c, p1Moves[i]), append(moves, p1Moves[i]), movesLeft-1)
//		if solved {
//			results = append(results, append(moves, p1Moves[i]))
//		}
//	}
//
//}
//
