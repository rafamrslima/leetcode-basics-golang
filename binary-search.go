func returnPosition(list []int, n int) int {
	left, right := 0, len(list)-1

	for left <= right {
		mid := (left + right) / 2
		if n < list[mid] {
			right = mid - 1
		} else if n > list[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
