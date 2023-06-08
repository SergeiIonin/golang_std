package main

import "fmt"

func main() {
	in1_0 := []int{1, 2, 2, 1}
	in2_0 := []int{2, 2}

	in1_1 := []int{4, 9, 5}
	in2_1 := []int{9, 4, 9, 8, 4}

	res0 := intersect(in1_0, in2_0)
	res1 := intersect(in1_1, in2_1)

	fmt.Println(res0)
	fmt.Println(res1)
}

func intersect(nums1 []int, nums2 []int) []int {

	res := make([]int, 0)

	hash := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]]++
	}

	for j := 0; j < len(nums2); j++ {
		v, ok := hash[nums2[j]]
		if ok {
			if v >= 1 {
				hash[nums2[j]]--
				res = append(res, nums2[j])
			}
		}
	}

	return res

}
