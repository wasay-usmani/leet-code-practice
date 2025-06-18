package trappingrainwater

func Trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	trapped := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				trapped += leftMax - height[left]
			}

			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				trapped += rightMax - height[right]
			}

			right--
		}
	}

	return trapped
}
