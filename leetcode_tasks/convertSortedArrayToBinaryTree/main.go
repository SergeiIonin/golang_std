package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// beats 53.9% RT and 74.83% Mem
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
		var left []int
		var right []int

		if size%2 != 0 {
			left = nums[:size/2]
			right = nums[size/2+1:]

			root = &TreeNode{
				Val:   nums[size/2],
				Left:  sortedArrayToBST(left),
				Right: sortedArrayToBST(right),
			}
		} else {
			left = nums[:size/2-1]
			right = nums[size/2:]

			root = &TreeNode{
				Val:   nums[size/2-1],
				Left:  sortedArrayToBST(left),
				Right: sortedArrayToBST(right),
			}
		}
		return root
	}
}
