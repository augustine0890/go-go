package main

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for idx := range nums {
		numsMap[nums[idx]] = true
	}

	longestStreak := 0
	for num := range numsMap {
		if numsMap[num - 1] {
			continue
		}
		
	}
	return longestStreak
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {

}
