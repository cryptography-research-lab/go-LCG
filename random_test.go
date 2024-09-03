package go_LCG

import "testing"

func TestLCG_Next(t *testing.T) {
	lcg := NewLCG()
	for i := 0; i < 100; i++ {
		println(lcg.Next())
	}
}
