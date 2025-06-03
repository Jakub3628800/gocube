package solver

import "testing"

func TestIdentity(t *testing.T) {
	cube := NewCubieCube()
	if !isSolved(&cube) {
		t.Fatalf("cube should start solved")
	}
}

func TestSolveNoop(t *testing.T) {
        sol, err := Solve("", 0)
        if err != nil {
                t.Fatalf("unexpected error: %v", err)
        }
        if len(sol) != 0 {
                t.Fatalf("expected empty solution")
        }
}

func TestCubieHelpers(t *testing.T) {
        c := NewCubieCube()
        if c.getTwist() != 0 || c.getFlip() != 0 {
                t.Fatalf("new cube should have zero twist and flip")
        }
        c.setTwist(123)
        if c.getTwist() != 123 {
                t.Fatalf("twist mismatch")
        }
        c.setFlip(5)
        if c.getFlip() != 5 {
                t.Fatalf("flip mismatch")
        }
        _ = c.cornerParity()
        _ = c.edgeParity()
        _ = c.getFRtoBR()
        c.setFRtoBR(0)
        other := NewCubieCube()
        c.multiply(&other)
}
