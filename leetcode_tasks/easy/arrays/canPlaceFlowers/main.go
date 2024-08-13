package main

import "fmt"

func main() {
	in := []int{1,0,0,0,1,0,0}
	res := canPlaceFlowers(in, 2)
	fmt.Printf("res = %v\n", res)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
    if n==0 {
        return true
    }

    count := 0

    if len(flowerbed)==1 {
        return n==1 && flowerbed[0]==0
    }

    sizeInit := len(flowerbed)

    cond := func(size int, rem int) bool {
        if size==1 {
            return rem==1
        }
        if size%2 == 0 {
           return size/2 >= rem
        } else {
            return size/2+1 >= rem
        }
    }
    
    if !cond(sizeInit, n) {
        return false
    }

    initIndex := 1

    if flowerbed[0]==0 && flowerbed[1]==0 {
		flowerbed[0]=1
		count++
        initIndex++    
	}

    for index:=initIndex ; index < sizeInit && count < n && cond(sizeInit-index, n-count) ; {
		step := 2

        if (flowerbed[index-1]==0 && flowerbed[index]==0) {
                if (index == len(flowerbed)-1) {
                        count++
                } else if (flowerbed[index+1]==0) {
                        flowerbed[index]=1
                        count++
                }
        } else if flowerbed[index]==0 {
				step = 1
		}
        index += step
    }
    return count==n
}