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

func cosineSimilarity(a_keys []int, a_values []float64, b_keys []int, b_values []float64) float64 {
	dictA := make(map[int]float64, len(a_keys))
	dictB := make(map[int]float64, len(b_keys))

	for i := range a_keys {
		dictA[a_keys[i]] = a_values[i]
	}
	for i := range b_keys {
		dictB[b_keys[i]] = b_values[i]
	}

	var dot float64

	for key, valA := range dictA {
		if valB, ok := dictB[key]; ok {
			dot += valA * valB
		}
	}

	var magA float64
	for _, value := range a_values {
		magA += value * value
	}

	var magB float64
	for _, value := range b_values {
		magB += value * value
	}

	magA = math.Sqrt(magA)
	magB = math.Sqrt(magB)

	if magA == 0 || magB == 0 {
		return 0
	}
	return dot / (magA * magB)
}
