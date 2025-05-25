package main

import (
	"cube/core"
	"cube/tables"
	"flag"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	// Initialize the logger for main package
	logger = log.New(os.Stdout, "main: ", log.LstdFlags)
}

func main() {
	generateTableFlag := flag.Bool("generateTable", false, "Generate and save the phase 1 pruning table.")
	flag.Parse()

	if *generateTableFlag {
		logger.Println("Generate table flag is set. Starting table generation process...")
		// The init functions in core package (InitMoveTables, InitP1Moves) are automatically called.
		generatedTable, err := tables.GeneratePhase1PruningTable(6) // Max depth of 6 for testing
		if err != nil {
			logger.Fatalf("Error generating Phase 1 pruning table: %v", err)
		}

		logger.Println("Phase 1 pruning table generated successfully.")
		err = tables.SavePhase1Table("phase1_depth6.gob", generatedTable)
		if err != nil {
			logger.Fatalf("Error saving Phase 1 pruning table: %v", err)
		}
		logger.Println("Phase 1 pruning table saved successfully to phase1_depth6.gob")
		logger.Println("Table generation complete. Exiting.")
		return // Exit after generating table
	}

	logger.Println("Normal execution. Attempting to load Phase 1 pruning table...")
	loadedTable, err := tables.LoadPhase1Table("phase1_depth6.gob")
	if err != nil {
		logger.Printf("Warning: Could not load Phase 1 pruning table 'phase1_depth6.gob': %v. Proceeding without table.", err)
		// Depending on application logic, may want to handle this more gracefully
		// or exit if the table is critical for normal operation.
	} else {
		logger.Println("Phase 1 pruning table 'phase1_depth6.gob' loaded successfully.")
		// You can now use loadedTable if needed for solver logic, e.g., pass it to a solver.
		// For now, we just log its successful load.
		// Example: check a value (ensure indices are valid and table is not nil)
		if loadedTable != nil {
			// These are example coordinates, replace with actual test coordinates if desired
			// For a solved cube: co=0, eo=0, us=494 (as per typical Kociemba coordinate definitions)
			// Note: The UDSlice coordinate for a solved cube might vary based on exact definition.
			// The Kociemba paper implies C(11,4)+C(10,3)+C(9,2)+C(8,1) = 330+120+36+8 = 494 for the 0th UD-slice combination.
			// Let's assume solved state is co=0, eo=0, us= (Value depends on UdsCoordinate output for solved cube)
			// For this example, let's just check the entry for a solved cube, which should be 0.
			initialCube := core.NewCube()
			co := core.CornerOrientationCoordinate(&initialCube)
			eo := core.EdgeOrientationCoordinate(&initialCube)
			us := core.UdsCoordinate(&initialCube)
			if co < 2187 && eo < 2048 && us < 495 { // Check bounds
				logger.Printf("Value from loaded table for solved state (CO:0, EO:0, US:%d): %d", us, loadedTable[co][eo][us])
			} else {
				logger.Printf("Coordinates for solved state (CO:%d, EO:%d, US:%d) are out of bounds for table lookup.", co, eo, us)
			}
		}
	}

	// Original application logic
	logger.Println("Proceeding with original cube operations...")
	c := core.NewCube()
	core.PrintCube(&c)
	core.PrintCoordinates(&c)

	c.Move("u", 2)
	c.Move("r", 1)
	core.PrintCube(&c)
	core.PrintCoordinates(&c)
	logger.Println("Cube operations complete.")
}
