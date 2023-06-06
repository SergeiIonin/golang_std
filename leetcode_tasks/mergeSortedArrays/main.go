package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	// nums1 := []int{2, 0}
	// m := 1
	// nums2 := []int{1}
	// n := 1

	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) > 0 {
		buf := append([]int{}, nums1[:m]...)
		k := 0 // for buf1
		j := 0 // for nums2
		for i := 0; i < len(nums1); i++ {
			if j < n {
				if k != m {
					head1 := buf[k]
					head2 := nums2[j]
					if head1 < head2 {
						nums1[i] = head1
						k++
					} else {
						nums1[i] = head2
						j++
					}
				} else {
					nums1[i] = nums2[j]
					j++
				}
			} else {
				nums1[i] = buf[k]
				k++
			}
		}
	}
}

// NB it may be more convenient to start from the end of the array (solution from leetcode):
// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	i, j, k := m-1, n-1, m+n-1
// 	for i >= 0 && j >= 0 {
// 		if nums1[i] > nums2[j] {
// 			nums1[k] = nums1[i]
// 			i--
// 		} else {
// 			nums1[k] = nums2[j]
// 			j--
// 		}
// 		k--
// 	}
// 	for j >= 0 {
// 		nums1[j] = nums2[j]
// 		j--
// 	}
// }
