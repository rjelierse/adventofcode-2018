package day03

import "testing"

func TestClaim_Occupies(t *testing.T) {
	c := Claim{1, 1, 3, 4, 4}
	o := c.Occupies()

	if len(o) != 16 {
		t.Errorf("Unexpected number of tiles. Expected 16, got %d.", len(o))
	}
}
