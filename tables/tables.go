package tables

import (
	"cube/core"
	"encoding/gob"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "tables: ", log.LstdFlags)
}

// indexed as [cor][eor][udslice] ~ 2,2 GB
type phase1Table [2187][2048][495]uint8

// init table with 255 (unvisited), save to a file if requested.
func Initphase1Table(save bool) (*phase1Table, error) {
	logger.Println("Initializing phase1 table with 255")
	var table phase1Table // Use var to allocate on the heap if it's too large for stack, though Go might do this anyway for large arrays.
	// Iterate and set all values to 255
	for i := range table {
		for j := range table[i] {
			for k := range table[i][j] {
				table[i][j][k] = 255
			}
		}
	}
	logger.Println("Finished initializing phase1 table with 255")

	if save {
		// Note: Saving a 2.2GB file of 255s might be slow and not very useful
		// unless it's a specific requirement for some external tool.
		// Consider if saving an unpopulated table is truly needed.
		logger.Println("Save flag is true, proceeding to save initialized table.")
		err := SavePhase1Table("phase1_initialized.gob", &table)
		if err != nil {
			logger.Printf("Error saving initialized phase1 table: %v", err)
			return &table, err
		}
		logger.Println("Initialized phase1 table saved.")
	}
	return &table, nil
}

// BFSState holds the cube state and its depth for the BFS queue.
type BFSState struct {
	cube  core.Cube
	depth uint8
}

// GeneratePhase1PruningTable creates the pruning table using BFS.
// maxSearchDepth is the maximum depth of moves stored in the table.
func GeneratePhase1PruningTable(maxSearchDepth uint8) (*phase1Table, error) {
	table, err := Initphase1Table(false) // Initialize table with 255s, don't save it yet.
	if err != nil {
		logger.Printf("Error initializing phase1 table for generation: %v", err)
		return nil, err
	}

	logger.Println("Starting Phase 1 Pruning Table generation...")

	initialCube := core.NewCube()
	co := core.CornerOrientationCoordinate(&initialCube) // Should be 0 for solved state
	eo := core.EdgeOrientationCoordinate(&initialCube)   // Should be 0 for solved state
	us := core.UdsCoordinate(&initialCube)               // Should be 494 for solved state (C(11,4)+C(10,3)+C(9,2)+C(8,1) based on typical Kociemba)

	// Ensure initial coordinates are within table bounds
	if co >= 2187 || eo >= 2048 || us >= 495 {
		logger.Fatalf("Initial coordinates out of bounds: co=%d, eo=%d, us=%d", co, eo, us)
		// os.Exit(1) or return error
	}
	
	if table[co][eo][us] == 255 { // Check if the solved state is unvisited (it should be)
		table[co][eo][us] = 0 // Depth of solved state is 0
	} else {
		// This case should ideally not happen if Initphase1Table correctly sets to 255
		logger.Printf("Warning: Solved state was already visited or not 255. table[%d][%d][%d] = %d", co, eo, us, table[co][eo][us])
		// If it's not 0, something is wrong. For safety, set to 0 if allowing continuation.
		table[co][eo][us] = 0 
	}


	queue := []BFSState{{cube: initialCube, depth: 0}}
	visitedCount := 1

	logger.Printf("Initial state: CO=%d, EO=%d, US=%d, Depth=0. Queue size: %d", co, eo, us, len(queue))


	processedStates := 0 // Counter for logging progress

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:] // Dequeue

		currentDepth := currentState.depth
		processedStates++

		if processedStates%100000 == 0 { // Log progress every 100,000 states processed
			logger.Printf("Processed %d states. Current depth: %d, Queue size: %d, Visited count: %d", processedStates, currentDepth, len(queue), visitedCount)
		}
		
		// If current state's depth is already maxSearchDepth, we don't explore its children
		// because their depth would be maxSearchDepth + 1, which is beyond what we want to store.
		if currentDepth >= maxSearchDepth {
			continue
		}

		for _, move := range core.P1Moves {
			newCube := currentState.cube // Create a copy
			newCube.Move(move.F, int(move.N))

			nco := core.CornerOrientationCoordinate(&newCube)
			neo := core.EdgeOrientationCoordinate(&newCube)
			nus := core.UdsCoordinate(&newCube)
			
			// Ensure new coordinates are within table bounds
			if nco >= 2187 || neo >= 2048 || nus >= 495 {
				logger.Printf("Warning: New coordinates out of bounds: nco=%d, neo=%d, nus=%d for move %v from depth %d. Skipping.", nco, neo, nus, move, currentDepth)
				continue 
			}

			newDepth := currentDepth + 1 // Children are one level deeper

			if table[nco][neo][nus] == 255 { // If unvisited
				table[nco][neo][nus] = newDepth
				queue = append(queue, BFSState{cube: newCube, depth: newDepth})
				visitedCount++
			}
		}
	}

	logger.Printf("Phase 1 Pruning Table generation complete. Visited states: %d. Processed states: %d", visitedCount, processedStates)
	return table, nil
}


func SavePhase1Table(filename string, table *phase1Table) error {

	logger.Println("Saving phase1 table")
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(&table); err != nil {
		return err
	}
	logger.Println("Phase1 table saved")
	return nil
}

func LoadPhase1Table(filename string) (*phase1Table, error) {
	logger.Println("Loading phase1 table")

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
	logger.Println("Loading phase1 table finished")
	return &table, nil
}

func Insertphase1TableItem(cube *core.Cube, lenMoves uint8, table *phase1Table) {
	cor := core.CornerOrientationCoordinate(cube)
	eor := core.EdgeOrientationCoordinate(cube)
	udslice := core.UdsCoordinate(cube)

	table[cor][eor][udslice] = lenMoves
}

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
