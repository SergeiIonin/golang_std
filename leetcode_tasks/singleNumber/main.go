package main

func main() {

}

func singleNumber(nums []int) int {
	sum := 0
	signMap := map[bool]int{true: 1, false: -1}
	hash := make(map[int]int, len(nums)/2)
	for i := 0; i < len(nums); i++ {
		_, ok := hash[nums[i]]
		hash[nums[i]] = i
		sum = sum + signMap[ok]*nums[i]
	}
	return sum
}

// the best(?) solution from leetcode is

// func singleNumber_(nums []int) int {
// 	res := 0
// 	for i := range nums {
// 		res = res ^ nums[i]
// 	}
// 	return res
// }

// that's because ^ (XOR) satisfies the following props:
// bit1 ^ bit2 = 1 if and only if ONLY ONE bit is 1
// 1) A ^ A == 0 (bc there's (1,0) pair of bits)
// 2) 0 ^ A = A
// 3) (A ^ B) ^ A == B
