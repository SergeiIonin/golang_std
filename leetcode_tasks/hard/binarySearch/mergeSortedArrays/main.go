package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 12, 14, 16}
	nums2 := []int{-2, -1, 0, 6, 8, 10, 18, 20, 22}
	res := make([]int, 0, len(nums1)+len(nums2))
	res1 := merge(nums1, nums2, res)
	res2 := merge(nums2, nums1, res)
	fmt.Println(res1)
	fmt.Println(res2)
}

func merge(nums1 []int, nums2 []int, res []int) []int {
	size1 := len(nums1)
	size2 := len(nums2)

	if size1 == 0 {
		return nums2
	} else if size2 == 0 {
		return nums1
	} else {
		if nums1[size1-1] <= nums2[0] {
			return append(nums1, nums2...)
		} else if nums2[size2-1] <= nums1[0] {
			return append(nums2, nums1...)
		} else {
			i := 0
			j := 0
			app1 := []int{}
			app2 := []int{}
			for i < size1 {
				elem1 := nums1[i]
				for j < size2 {
					if nums2[j] <= elem1 {
						if len(app1) != 0 {
							res = append(res, app1...)
							app1 = []int{}
						}
						app2 = append(app2, nums2[j])
						j++
					} else if nums2[j] > elem1 {
						if len(app2) != 0 {
							res = append(res, app2...)
							app2 = []int{}
						}
						app1 = append(app1, nums1[i])
						i++
						break
					}
				}
				if j == size2 {
					res = append(res, app2...)
					return append(res, nums1[i:]...)
				}
			}
			res = append(res, app1...)
			return append(res, nums2[j:]...)
		}
	}
}
