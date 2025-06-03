package main

import (
	"cube/solver"
	"fmt"
)

func main() {
	scramble := "BBURUDBFUFFFRRFUUFLULUFUDLRRDBBDBDBLUDDFLLRRBRLLLBRDDF"
	moves, err := solver.Solve(scramble, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("Solution:", moves)
}
