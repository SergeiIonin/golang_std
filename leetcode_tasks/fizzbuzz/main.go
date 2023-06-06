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

// This solution is less performant:
// func fizzBuzz(n int) []string {

//     if n == 0 {
//         return []string{}
//     }

//     res := make([]string, n)

//     i := 2 // 3
//     j := 4 // 5

//     for l:=0; l<n; l++ {
//         for i<n {
//             res[i] += "Fizz"
// 			i+=3
//         }
// 		for j<n {
// 			res[j] += "Buzz"
// 			j+=5
// 		}
//         if res[l] == "" {
//             res[l] = fmt.Sprintf("%d", l+1)
//         }
//     }

//     return res
// }
