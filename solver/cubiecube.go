package solver

const (
	URF = iota
	UFL
	ULB
	UBR
	DFR
	DLF
	DBL
	DRB
)

const (
	UR = iota
	UF
	UL
	UB
	DR
	DF
	DL
	DB
	FR
	FL
	BL
	BR
)

type CubieCube struct {
	Cp [8]int
	Co [8]int
	Ep [12]int
	Eo [12]int
}

func NewCubieCube() CubieCube {
	return CubieCube{
		Cp: [8]int{URF, UFL, ULB, UBR, DFR, DLF, DBL, DRB},
		Co: [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		Ep: [12]int{UR, UF, UL, UB, DR, DF, DL, DB, FR, FL, BL, BR},
		Eo: [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func (c *CubieCube) cornerMultiply(b *CubieCube) {
	var cp [8]int
	var co [8]int
	for i := 0; i < 8; i++ {
		cp[i] = c.Cp[b.Cp[i]]
		oriA := c.Co[b.Cp[i]]
		oriB := b.Co[i]
		co[i] = (oriA + oriB) % 3
	}
	c.Cp = cp
	c.Co = co
}

func (c *CubieCube) edgeMultiply(b *CubieCube) {
	var ep [12]int
	var eo [12]int
	for i := 0; i < 12; i++ {
		ep[i] = c.Ep[b.Ep[i]]
		eo[i] = (b.Eo[i] + c.Eo[ep[i]]) % 2
	}
	c.Ep = ep
	c.Eo = eo
}

func (c *CubieCube) multiply(b *CubieCube) {
	c.cornerMultiply(b)
	c.edgeMultiply(b)
}

func rotateLeft(arr []int, l, r int) {
	t := arr[l]
	copy(arr[l:r], arr[l+1:r+1])
	arr[r] = t
}

func rotateRight(arr []int, l, r int) {
	t := arr[r]
	for i := r; i > l; i-- {
		arr[i] = arr[i-1]
	}
	arr[l] = t
}

func Cnk(n, k int) int {
	if n < k {
		return 0
	}
	if k > n/2 {
		k = n - k
	}
	s := 1
	for i, j := n, 1; i > n-k; i, j = i-1, j+1 {
		s = s * i / j
	}
	return s
}

func (c *CubieCube) getTwist() int {
	ret := 0
	for i := 0; i < 7; i++ {
		ret = 3*ret + c.Co[i]
	}
	return ret
}

func (c *CubieCube) setTwist(twist int) {
	twistParity := 0
	for i := 6; i >= 0; i-- {
		c.Co[i] = twist % 3
		twistParity += c.Co[i]
		twist /= 3
	}
	c.Co[7] = (3 - twistParity%3) % 3
}

func (c *CubieCube) getFlip() int {
	ret := 0
	for i := 0; i < 11; i++ {
		ret = 2*ret + c.Eo[i]
	}
	return ret
}

func (c *CubieCube) setFlip(flip int) {
	flipParity := 0
	for i := 10; i >= 0; i-- {
		c.Eo[i] = flip % 2
		flipParity += c.Eo[i]
		flip /= 2
	}
	c.Eo[11] = (2 - flipParity%2) % 2
}

func (c *CubieCube) cornerParity() int {
	s := 0
	for i := 7; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if c.Cp[j] > c.Cp[i] {
				s++
			}
		}
	}
	return s % 2
}

func (c *CubieCube) edgeParity() int {
	s := 0
	for i := 11; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if c.Ep[j] > c.Ep[i] {
				s++
			}
		}
	}
	return s % 2
}

func (c *CubieCube) getFRtoBR() int {
	a := 0
	x := 0
	edge4 := [4]int{}
	for j := 11; j >= 0; j-- {
		if FR <= c.Ep[j] && c.Ep[j] <= BR {
			a += Cnk(11-j, x+1)
			edge4[3-x] = c.Ep[j]
			x++
		}
	}
	b := 0
	for j := 3; j > 0; j-- {
		k := 0
		for edge4[j] != j+8 {
			rotateLeft(edge4[:], 0, j)
			k++
		}
		b = (j+1)*b + k
	}
	return 24*a + b
}

func (c *CubieCube) setFRtoBR(idx int) {
	sliceEdge := []int{FR, FL, BL, BR}
	otherEdge := []int{UR, UF, UL, UB, DR, DF, DL, DB}
	b := idx % 24
	a := idx / 24
	for i := 0; i < 12; i++ {
		c.Ep[i] = DB
	}
	for j := 1; j < 4; j++ {
		k := b % (j + 1)
		b /= j + 1
		for k > 0 {
			rotateRight(sliceEdge, 0, j)
			k--
		}
	}
	x := 3
	for j := 0; j < 12; j++ {
		if a-Cnk(11-j, x+1) >= 0 {
			c.Ep[j] = sliceEdge[3-x]
			a -= Cnk(11-j, x+1)
			x--
		}
	}
	x = 0
	for j := 0; j < 12; j++ {
		if c.Ep[j] == DB {
			c.Ep[j] = otherEdge[x]
			x++
		}
	}
}

var moveCube = [6]CubieCube{
	{
		Cp: [8]int{UBR, URF, UFL, ULB, DFR, DLF, DBL, DRB},
		Co: [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		Ep: [12]int{UB, UR, UF, UL, DR, DF, DL, DB, FR, FL, BL, BR},
		Eo: [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		Cp: [8]int{DFR, UFL, ULB, URF, DRB, DLF, DBL, UBR},
		Co: [8]int{2, 0, 0, 1, 1, 0, 0, 2},
		Ep: [12]int{FR, UF, UL, UB, BR, DF, DL, DB, DR, FL, BL, UR},
		Eo: [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		Cp: [8]int{UFL, DLF, ULB, UBR, URF, DFR, DBL, DRB},
		Co: [8]int{1, 2, 0, 0, 2, 1, 0, 0},
		Ep: [12]int{UR, FL, UL, UB, DR, FR, DL, DB, UF, DF, BL, BR},
		Eo: [12]int{0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0},
	},
	{
		Cp: [8]int{URF, UFL, ULB, UBR, DLF, DBL, DRB, DFR},
		Co: [8]int{0, 0, 0, 0, 0, 0, 0, 0},
		Ep: [12]int{UR, UF, UL, UB, DF, DL, DB, DR, FR, FL, BL, BR},
		Eo: [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		Cp: [8]int{URF, ULB, DBL, UBR, DFR, UFL, DLF, DRB},
		Co: [8]int{0, 1, 2, 0, 0, 2, 1, 0},
		Ep: [12]int{UR, UF, BL, UB, DR, DF, FL, DB, FR, UL, DL, BR},
		Eo: [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	{
		Cp: [8]int{URF, UFL, UBR, DRB, DFR, DLF, ULB, DBL},
		Co: [8]int{0, 0, 1, 2, 0, 0, 2, 1},
		Ep: [12]int{UR, UF, UL, BR, DR, DF, DL, BL, FR, FL, UB, DB},
		Eo: [12]int{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
	},
}

func (c *CubieCube) Move(m int) {
	orig := *c
	for i := 0; i < 8; i++ {
		idx := moveCube[m].Cp[i]
		c.Cp[i] = orig.Cp[idx]
		c.Co[i] = (orig.Co[idx] + moveCube[m].Co[i]) % 3
	}
	for i := 0; i < 12; i++ {
		idx := moveCube[m].Ep[i]
		c.Ep[i] = orig.Ep[idx]
		c.Eo[i] = (orig.Eo[idx] + moveCube[m].Eo[i]) % 2
	}
}
