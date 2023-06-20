package main

import (
	"fmt"
	"testing"
)

func TestMinWindow(t *testing.T) {

	inputs := []TestInputAndResult{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
			"[1] contains in substr of size 4",
		},
		{
			"a",
			"a",
			"a",
			"[2] both strings are the same",
		},
		{
			"a",
			"aa",
			"",
			"[3] second is bigger",
		},
		{
			"a",
			"b",
			"",
			"[4] two different single runes",
		},
		{
			"ab",
			"a",
			"a",
			"[5] ab and a",
		},
		{
			"ab",
			"b",
			"b",
			"[6] ab and b",
		},
		{
			"bdab",
			"ab",
			"ab",
			"[7] bdab and ab",
		},
		{
			"abc",
			"ac",
			"abc",
			"[8] abc and ac",
		},
	}
	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			got := minWindow(input.s1, input.s2)
			if got != input.res {
				fmt.Println("test ", input.name)
				fmt.Println("got = ", got)
				fmt.Println("expected = ", input.res)
				t.Error("test failed")
			}
		})
	}

}

type TestInputAndResult struct {
	s1   string
	s2   string
	res  string
	name string
}
