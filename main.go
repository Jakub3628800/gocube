package main

import (
	"cube/core"
	"fmt"
)

func main() {
	core.InitMoveTables()
	c := core.NewCube()
	c.Print()
	c.Move("u", 2)
	c.Move("r", 1)
	c.Print()
	c.PrintCoordinates()
	core.InitP1Moves()

	moves, solved := core.SolvePhase1(c, []core.Move{}, 1)
	fmt.Println(moves, solved)
}
