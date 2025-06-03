package solver

import "errors"

var moves = []int{0, 1, 2, 3, 4, 5}

func Solve(state string, depth int) ([]int, error) {
	// parse state string to cube
	cube := NewCubieCube()
	// TODO: parse state string
	// fallback: return error
	if len(state) != 0 {
		return nil, errors.New("parsing not implemented")
	}

	solution := []int{}
	if idaStar(&cube, 0, depth, &solution) {
		return solution, nil
	}
	return nil, errors.New("no solution")
}

func idaStar(c *CubieCube, depth, maxDepth int, path *[]int) bool {
	if depth == maxDepth {
		if isSolved(c) {
			return true
		}
		return false
	}
	for _, m := range moves {
		next := *c
		next.Move(m)
		*path = append(*path, m)
		if idaStar(&next, depth+1, maxDepth, path) {
			return true
		}
		*path = (*path)[:len(*path)-1]
	}
	return false
}

func isSolved(c *CubieCube) bool {
	for i := 0; i < 8; i++ {
		if c.Cp[i] != i || c.Co[i] != 0 {
			return false
		}
	}
	for i := 0; i < 12; i++ {
		if c.Ep[i] != i || c.Eo[i] != 0 {
			return false
		}
	}
	return true
}
