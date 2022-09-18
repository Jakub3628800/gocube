package main

type cube struct {
	cornerPermutation [8]int
	cornerOrientation [8]int
	edgePermutation   [8]int
	edgeOrientation   [8]int
}

// Return hash of cube object.
func (c *cube) hash() string {
	return "s"
}

// Is the cube solved?
func (c *cube) solved() bool {
	return true
}

// Apply move to cube
func (c *cube) move(moveStr string) {
}

func main() {

}
