package tables

import (
	"encoding/gob"
	"os"
)

// indexed as [cor][eor][udslice] ~ 2,2 GB
type phase1Table [2187][2048][495]uint8

// init table with zeros, save to a file.
func Initphase1Table() error {
	table := phase1Table{}
	err := SavePhase1Table("phase1.gob", &table)
	if err != nil {
		return err
	}
	return nil
}

func SavePhase1Table(filename string, table *phase1Table) error {

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(&table); err != nil {
		return err
	}
	return nil
}

func LoadPhase1Table(filename string) (*phase1Table, error) {

	table := phase1Table{}
	file, err := os.Open(filename)
	if err != nil {
		return &table, err
	}
	defer file.Close()
	//
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&table); err != nil {
		return &table, err
	}
	return &table, nil
}

//func insertphase1TableItem(item MoveItem, lenMoves uint8, table *phase1Table) {
//	cor := item.cube.CornerOrientationCoordinate()
//	eor := item.cube.EdgeOrientationCoordinate()
//	udslice := item.cube.UdsCoordinate()
//
//	table[cor][eor][udslice] = lenMoves
//
//}
//
//type MoveItem struct {
//	cube             core.Cube
//	lastMove         core.Move
//	previousLastMove core.Move
//	movesLeft        int
//}
//
//func allValidMoves() []core.Move {
//	return []core.Move{
//		core.Move{"r", 1},
//		core.Move{"r", 2},
//		core.Move{"r", 3},
//
//		core.Move{"l", 1},
//		core.Move{"l", 2},
//		core.Move{"l", 3},
//
//		core.Move{"u", 1},
//		core.Move{"u", 2},
//		core.Move{"u", 3},
//
//		core.Move{"d", 1},
//		core.Move{"d", 2},
//		core.Move{"d", 3},
//
//		core.Move{"f", 1},
//		core.Move{"f", 2},
//		core.Move{"f", 3},
//
//		core.Move{"b", 1},
//		core.Move{"b", 2},
//		core.Move{"b", 3},
//	}
//}
//
//func oppositeMove(move string) string {
//	switch {
//	case move == "r":
//		return "l"
//	case move == "f":
//		return "b"
//	case move == "u":
//		return "d"
//	case move == "l":
//		return "r"
//	case move == "b":
//		return "f"
//	case move == "d":
//		return "u"
//	default:
//		fmt.Println("It's after noon")
//	}
//	return ""
//}
//
//func createTablePhase1() {
//
//	movesLenght := 4 // How deep do we search
//	moves := allValidMoves()
//
//	solvedCube := core.NewCube()
//	moveQueue := []MoveItem{
//		MoveItem{
//			cube:             core.CreateCopyWithmove(solvedCube, core.Move{"r", 1}),
//			lastMove:         core.Move{"r", 1},
//			previousLastMove: core.Move{"r", 0},
//			movesLeft:        movesLenght,
//		},
//		MoveItem{
//			cube:             core.CreateCopyWithmove(solvedCube, core.Move{"l", 1}),
//			lastMove:         core.Move{"l", 1},
//			previousLastMove: core.Move{"l", 0},
//			movesLeft:        movesLenght,
//		},
//		MoveItem{
//			cube:             core.CreateCopyWithmove(solvedCube, core.Move{"f", 1}),
//			lastMove:         core.Move{"f", 1},
//			previousLastMove: core.Move{"f", 0},
//			movesLeft:        movesLenght,
//		},
//		MoveItem{
//			cube:             core.CreateCopyWithmove(solvedCube, core.Move{"b", 1}),
//			lastMove:         core.Move{"b", 1},
//			previousLastMove: core.Move{"b", 0},
//			movesLeft:        movesLenght,
//		},
//	}
//	var currentMove core.Move
//	var c core.Cube
//	var previous core.Move
//	var mi MoveItem
//	var moveIt MoveItem
//	var i int
//
//	for len(moveQueue) > 0 { // while BFS search queue is not empty
//
//		moveIt = moveQueue[0] // pop an item
//		moveQueue = moveQueue[1:]
//
//		for i = 0; i < len(moves); i++ {
//			currentMove = moves[i]
//			if currentMove.F == moveIt.lastMove.F || currentMove.F == oppositeMove(currentMove.F) {
//				continue
//			}
//			c = core.CreateCopyWithmove(moveIt.cube, currentMove)
//			previous = moveIt.lastMove
//			mi = MoveItem{cube: c, lastMove: currentMove, previousLastMove: previous, movesLeft: moveIt.movesLeft - 1}
//			if mi.movesLeft != 0 {
//				moveQueue = append(moveQueue, mi)
//			}
//		}
//		//fmt.Println(len(moveQueue))
//
//	}
//
//}

//func phase1BF(c *core.Cube, lastMove, previousLastMove, movesToSolve int, table *phase1Table) {
//
//	results := [][]Move{}
//	for i := 0; i < len(p1Moves); i++ {
//		moves, solved = SolvePhase1(createCopyWithMove(c, p1Moves[i]), append(moves, p1Moves[i]), movesLeft-1)
//		if solved {
//			results = append(results, append(moves, p1Moves[i]))
//		}
//	}
//
//}
