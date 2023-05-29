package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	dest := make(chan int)

	/*wgMain := sync.WaitGroup{}
	wgMain.Add(2)*/

	size := 10

	go func() {
		//defer wgMain.Done()
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		//close(ch1) // NB!
	}()

	go func() {
		//defer wgMain.Done()
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		//close(ch2)
	}()

	/*go func() {
		count := 0
		for {
			if count == size {
				break
			}
			select {
			case i := <-ch1:
				//fmt.Println("captured from ch1 value ", i)
				dest <- i
			case j := <-ch2:
				//fmt.Println("captured from ch2 value ", j)
				dest <- j
			}
			count++
		}
		close(ch1)
		close(ch2)
		close(dest)
	}()*/

	go func() {
		defer func() {
			close(ch1)
			close(ch2)
			close(dest)
		}()
		merge(&ch1, &ch2, &dest, size)
	}()

	for k := range dest {
		fmt.Println("received value dest = ", k)
		//fmt.Println("------------------", k)
	}
}

func merge(from1, from2 *chan int, dest *chan int, size int) {
	count := 0
	for {
		if count == size {
			break
		}
		select {
		case i := <-*from1:
			*dest <- i
		case j := <-*from2:
			*dest <- j
		}
		count++
	}
}

func merge1(from1, from2 chan int) chan int {
	dest := make(chan int)
	chan1Closed := false
	chan2Closed := false
	for {
		select {
		case i, open := <-from1:
			if !open {
				chan1Closed = true
				break
			}
			dest <- i
		case j, open := <-from2:
			if !open {
				chan2Closed = true
				break
			}
			dest <- j
		}
		if chan1Closed && chan2Closed {
			close(dest)
			break
		}
	}
	return dest
}
func merge2(from1, from2 chan int) chan int {
	dest := make(chan int)
	for from1 != nil && from2 != nil {
		select {
		case i, open := <-from1:
			if !open {
				from1 = nil
				break
			}
			dest <- i
		case j, open := <-from2:
			if !open {
				from2 = nil
				break
			}
			dest <- j
		}
	}
	close(dest)
	return dest
}

/*
func merge(from1, from2 chan int) chan int {

	dest := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		push(from1, dest)
	}()

	go func() {
		defer wg.Done()
		push(from2, dest)
	}()

	wg.Wait()

	return dest
}*/

func push(from, to chan int) {
	for i := range from {
		to <- i
	}
}
