package main

import "fmt"

func main() {
	fmt.Println(fizzBuzz(15))
	fmt.Println(fizzBuzz(30))
}

func fizzBuzz(n int) []string {

	if n == 0 {
		return []string{}
	}

	res := make([]string, n)

	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 {
			res[i] += "Fizz"
		}
		if (i+1)%5 == 0 {
			res[i] += "Bazz" // NB use += for strings
		}
		if res[i] == "" {
			res[i] = fmt.Sprintf("%d", i+1)
		}
	}

	return res
}
