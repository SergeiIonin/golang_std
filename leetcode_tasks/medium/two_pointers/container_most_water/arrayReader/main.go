package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("resources/input")
	if err != nil {
		fmt.Println("error = ", err.Error())
	}
	r := bufio.NewReader(f)
	ints := make([]int, 0, 10000)

	for {
		b, err := r.ReadBytes(byte(','))
		if err != nil {
			fmt.Println("end of file:", err.Error())
			break
		}
		bNew := bytes.TrimSuffix(b, []byte(","))
		//copy(bNew, bytes.TrimSuffix(b, []byte(",")))
		s := string(bNew)
		int, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err.Error())
		}
		ints = append(ints, int)
	}
	fmt.Println(ints)

}
