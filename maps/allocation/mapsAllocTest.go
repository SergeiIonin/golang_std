package main

import (
	"fmt"
	"runtime"
)

func main() {

	bytesArr := [128]byte(randBytesInit())

	n := 1_000_000
	m := make(map[int][128]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = bytesArr
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)

}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Mb\n", m.Alloc/1048576)
}

func randBytesInit() []byte {
	var bytes []byte
	var b byte
	for i := 0; i < 128; i++ {
		b = byte(i)
		bytes = append(bytes, b)
	}
	return bytes
}
