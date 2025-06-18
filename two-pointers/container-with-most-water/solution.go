package containerwithmostwater

// MaxArea returns the maximum area of water a container can store.
func MaxArea(height []int) int {
	max := 0
	var left, right int = 0, len(height) - 1
	for left < right {
		if height[left] >= height[right] {
			if height[right]*(right-left) > max {
				max = height[right] * (right - left)
			}

			right--
		}

		if height[left] < height[right] {
			if height[left]*(right-left) > max {
				max = height[left] * (right - left)
			}

			left++
		}
	}

	return max
}
