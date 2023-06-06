package main

import "fmt"

func main() {

}

func fizzBuzz(n int) []string {

	if n == 0 {
		return []string{}
	}

	res := make([]string, n)

	i := 2  // 3
	j := 4  // 5
	k := 14 // 15

	for l := 0; l < n; l++ {
		if k < n {
			res[k] = "FizzBuzz"
			k += 15
		}
		if j < n {
			if res[j] == "" {
				res[j] = "Buzz"
			}
			j += 5
		}
		if i < n {
			if res[i] == "" {
				res[i] = "Fizz"
			}
			i += 3
		}
		if res[l] == "" {
			res[l] = fmt.Sprintf("%d", l+1)
		}
	}

	return res
}
