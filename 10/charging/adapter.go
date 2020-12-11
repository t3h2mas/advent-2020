package charging

import (
	"sort"
)

func ChainVoltageDifferences(adapters []int) int {
	sort.Ints(adapters)

	// handle built in voltage by adding 1 to the differences of 3 group
	differenceMap := map[int]int{3: 1}
	// handle the voltage difference from the initial device (0)
	differenceMap[adapters[0]]++
	differences := voltageDifferences(adapters, differenceMap)

	return differences[1] * differences[3]
}

func voltageDifferences(adapters []int, differences map[int]int) map[int]int {
	if len(adapters) <= 1 {
		return differences
	}

	adapterVoltage := adapters[0]
	diff := adapters[1] - adapterVoltage

	differences[diff]++

	return voltageDifferences(adapters[1:], differences)
}

func ChainPossibilities(adapters []int) int {
	maxVoltage := max(adapters)
	available := make([]int, maxVoltage+1)

	for _, voltage := range adapters {
		available[voltage] = 1
	}

	dp := make([]int, len(available))
	dp[0] = 1

	for i := 1; i < len(available); i++ {
		if available[i] == 0 {
			continue
		}
		for j := 1; j < 4; j++ {
			if i-j < 0 {
				continue
			}
			// skip voltages that don't match an adapter
			if available[i] == 0 {
				dp[i] = 0
			} else {
				dp[i] += dp[(i - j)]
			}
		}
	}

	return dp[len(available)-1]
}

func max(nums []int) int {
	m := nums[0]

	for _, n := range nums {
		if n > m {
			m = n
		}
	}

	return m
}
