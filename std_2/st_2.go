package main

import "fmt"

func main() {

	fmt.Println("---ARRAYS---")

	var a [5]int
	fmt.Println("empty array", a)

	a[3] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)

	fmt.Println("---SLICES---")
	s := make([]string, 3)
	fmt.Println("empty slice:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set slice:", s)
	fmt.Println("get from slice:", s[2])
	fmt.Println("len of slice:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("slice updated:", s)

	c := make([]string, len(s))

	copy(c, s)
	fmt.Println("copied:", c)

	l := s[2:5]
	fmt.Println("s[2:5]", l)

	l = s[2:] // note, = instead of :=
	fmt.Println("s[2:]", l)

	l = s[:5]
	fmt.Println("s[:5]", l)

	t := []string{"g", "h", "i"}
	fmt.Println("slice created and initialized", t)

	twoDslice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoDslice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoDslice[i][j] = i + j
		}
	}
	fmt.Println("twoDslice = ", twoDslice)

	fmt.Println("---MAPS---")

	m := make(map[string]int)

	m["k1"] = 12
	m["k2"] = 24

	fmt.Println("map = ", m)

	v1 := m["k1"] // note get syntax is the same
	fmt.Println("v1 = ", v1)

	v3 := m["k3"]
	fmt.Println("v3 = ", v3)
	fmt.Println("len of map", len(m))

	delete(m, "k2")
	fmt.Println("map = ", m)

	_, prs := m["k2"]
	fmt.Println("map contains k2 = ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map created and initialized", n)

	fmt.Println("---RANGE---")
	//range iterates over elements in a variety of data structures
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	for i, _ := range nums { // we can access index and value in 'for'
		if i == 3 {
			fmt.Println("index of 3 is", i)
		}
	}
	fmt.Println("sum = ", sum)

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	for k := range kvs {
		fmt.Println("key = ", k)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println("---FUNCTIONS---")

	plusRes := plus(1, 2)
	fmt.Println("plusRes = ", plusRes)
	plusManyRes := plusMany(1, 2, 3)
	fmt.Println("plusManyRes = ", plusManyRes)

	fmt.Println("---MULTIPLE RETURN VALUES---")
	//Go has built-in support for multiple return values. This feature is used often in idiomatic Go, for example
	//to return both result and error values from a function.
	aMRV, bMRV := vals()
	_, bMRVonly := vals()
	fmt.Println(aMRV, bMRV)
	fmt.Println(bMRVonly)

	fmt.Println("---CLOSURES---")

	nextInt := intSeq() // i is bounded to the closure
	fmt.Println("nextInt = ", nextInt())
	fmt.Println("nextInt = ", nextInt())
	fmt.Println("nextInt = ", nextInt())

	newInt := intSeq()
	fmt.Println("newInt = ", newInt())

}

func plus(a int, b int) int {
	return a + b
}

func plusMany(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

// This function intSeq returns another function, which we define anonymously in the body of intSeq.
// The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
