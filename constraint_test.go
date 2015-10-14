package typewriter

import "testing"

func TestTryType(t *testing.T) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			typ := Type{
				Comparable: i == 0,
				Numeric:    i == 1,
				Ordered:    i == 2,
			}
			c := Constraint{
				Comparable: j == 0,
				Numeric:    j == 1,
				Ordered:    j == 2,
			}

			err := c.CheckCompatibility(typ)
			should := i == j

			if should != (err == nil) {
				t.Errorf("TryType is incorrect when for Type %v on Constraint %v; should be %v", typ, c, should)
			}
		}
	}
}
