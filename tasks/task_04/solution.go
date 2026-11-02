package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	count := len(nums) - 1
	sum := int64(0)
	minimum := nums[1] - nums[0]
	maximum := nums[1] - nums[0]
	for i := 1; i < count+1; i++ {
		diff := nums[i] - nums[i-1]
		sum += diff
		minimum = min(minimum, diff)
		maximum = max(maximum, diff)
	}
	return Stats{count, sum, minimum, maximum}
}
