func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		_, ok := numsMap[complement]

		if numsMap[complement] != i && ok {
			return []int{i, numsMap[complement]}
		}

		numsMap[nums[i]] = i
	}
	return nil
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		} else {
			seen[num] = true
		}
	}
	return false
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := make(map[rune]int)

	for _, ch := range s {
		dict[ch]++
	}

	for _, ch := range t {
		count, ok := dict[ch]
		if !ok || count == 0 {
			return false
		}
		dict[ch]--
	}
	return true
}
