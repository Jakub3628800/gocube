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
