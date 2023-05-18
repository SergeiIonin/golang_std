package main

import "fmt"

func main() {

	s1 := make([]int, 3, 6)

	s1[0] = 1
	s1[1] = 2
	s1[2] = 3

	s2 := s1[1:3]

	s2 = append(s2, 4)

	fmt.Println("s1 = ", s1) // [1 2 3]
	fmt.Println("s2 = ", s2) // [2 3 4]

	s1 = append(s1, 5)

	fmt.Println("s1 = ", s1) // [1 2 3 5]
	fmt.Println("s2 = ", s2) // [2 3 5]

}
