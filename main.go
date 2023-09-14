package main

import (
	"cube/core"
)

func main() {
	core.InitMoveTables()
	c := core.NewCube()
	core.PrintCube(&c)
	core.PrintCoordinates(&c)
	c.Move("u", 2)
	c.Move("r", 1)
	core.PrintCube(&c)
	core.PrintCoordinates(&c)

}
