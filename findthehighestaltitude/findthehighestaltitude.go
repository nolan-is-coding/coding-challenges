package findthehighestaltitude

func max(slice []int) int {
	m := slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i] > m {
			m = slice[i]
		}
	}
	return m
}

// gain = [-5,1,5,0,-6]
// altitudes [0,-5,-4,1,1,-5]
func FindTheHighestAltitude(gain []int) int {
	gain = append([]int{0}, gain...)
	res := make([]int, len(gain))
	for i := 1; i < len(gain); i++ {
		res[i] = res[i-1] + gain[i]
	}
	r := max(res)
	return r
}
