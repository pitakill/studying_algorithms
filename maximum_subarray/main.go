package maximum_subarray

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for _, n := range nums[1:] {
		if n < n+sum {
			sum += n
		} else {
			sum = n
		}

		if max < sum {
			max = sum
		}
	}

	return max
}
