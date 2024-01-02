package kata

func maxArea(height []int) int {
	area := 0
	i, j := 0, len(height)-1

	for i < j {
		h := min(height[i], height[j])
		w := j - i
		area = max(area, h*w)

		if height[i] < height[j] {
			i++
		} else {
			j++
		}
	}

	return area
}
