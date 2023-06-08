package main

func main() {

}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int, len(nums)) // we can also use map[int]struct{} as in the most performant solution on leetcode
	for i := 0; i < len(nums); i++ {
		_, ok := hash[nums[i]]
		if ok {
			return true
		} else {
			hash[nums[i]] = i
		}
	}
	return false
}
