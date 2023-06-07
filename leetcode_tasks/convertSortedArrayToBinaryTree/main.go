package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// beats 100% RT and 74.83% Mem
func sortedArrayToBST(nums []int) *TreeNode {
	size := len(nums)

	if size == 0 {
		return nil
	} else if size == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	} else if size == 2 {
		return &TreeNode{
			Val:  nums[0],
			Left: nil,
			Right: &TreeNode{
				Val:   nums[1],
				Left:  nil,
				Right: nil,
			},
		}
	} else {
		var root *TreeNode
		var middle int
		var left []int
		var right []int

		if size%2 != 0 {
			middle = nums[size/2]
			left = nums[:size/2]
			right = nums[size/2+1:]
		} else {
			middle = nums[size/2-1]
			left = nums[:size/2-1]
			right = nums[size/2:]
		}
		root = &TreeNode{
			Val:   middle,
			Left:  sortedArrayToBST(left),
			Right: sortedArrayToBST(right),
		}
		return root
	}
}

// less verbose solution from leetcode:
// func sortedArrayToBST(nums []int) *TreeNode {
//     if len(nums) == 0{
//         return nil
//     }
//     mid_node := len(nums) / 2

//     tree := &TreeNode{Val:nums[mid_node], Left:sortedArrayToBST(nums[:mid_node]), Right:sortedArrayToBST(nums[mid_node+1:])}
//     return tree

// }
