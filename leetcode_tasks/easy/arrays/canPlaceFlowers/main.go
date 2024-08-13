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

    for index:=0 ; index < sizeInit && count < n && cond(sizeInit-index, n-count) ; {
        fmt.Printf("index = %v, count = %d \n", index, count)
		step := 1
        if (index==0) {
            if flowerbed[0]==0 && flowerbed[1]==0 {
                flowerbed[0]=1
                count++
                step += 1            
            }
        } else {
            if (flowerbed[index-1]==0 && flowerbed[index]==0) {
                if (index == len(flowerbed)-1) {
                        count++
                        step += 1
                } else if (flowerbed[index+1]==0) {
                        flowerbed[index]=1
                        count++
                        step += 1
                }
            }
        }
        index += step
		c := cond(sizeInit-index, n-count)
		fmt.Printf("count (iter end) = %d, c = %v \n", count, c)
    }
    return count==n
}