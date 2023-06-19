package main

import (
	"fmt"
	"testing"
)

func TestTimeBasedKVStore(t *testing.T) {

	inputs := []TestInputAndResult{
		{
			"love",
			"",
			5,
			"ts 5, should return empty str",
		},
		{
			"love",
			"high",
			10,
			"ts 10 should return high",
		},
		{
			"love",
			"high",
			15,
			"ts 15 should return high",
		},
		{
			"love",
			"low",
			20,
			"ts 20 should return low",
		},
		{
			"love",
			"low",
			25,
			"ts 25 should return low",
		},
	}
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	for _, input := range inputs {
		t.Run(input.name, func(t *testing.T) {
			got := tm.Get(input.key, input.ts)
			if got != input.value {
				fmt.Println("test ", input.name)
				fmt.Println("got = ", got)
				t.Error("test failed")
			}
		})
	}

}

type TestInputAndResult struct {
	key   string
	value string
	ts    int
	name  string
}
