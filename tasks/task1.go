package tasks

// FindTheMinNum returns the minimal number and the number appears frequency.
func FindTheMinNum(arr []int) (min, frequency int) {
	m := make(map[int]int)
	min = 100000

	for _, num := range arr {
		if num < min {
			min = num
		}
		m[num]++
	}

	return min, m[min]
}
