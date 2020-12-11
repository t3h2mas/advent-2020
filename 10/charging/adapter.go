package charging

import (
	"sort"

	"github.com/t3h2mas/advent-2020/types"
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
	sort.Ints(adapters)
	// add the outlet device voltage of 0
	adapters = append([]int{0}, adapters...)

	// add the built in adapter voltage of 3 + the max voltage
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	available := types.NewIntSet()

	for _, voltage := range adapters {
		available.Add(voltage)
	}

	dp := make([]int, len(adapters))
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < len(adapters); i++ {
		for j := 1; j < 4; j++ {
			if available.Has(adapters[i] - j) {
				dp[i] += dp[(i - j)]
			}
		}

	}

	//return dp[len(available)-1]
	return max(dp)
}

func max(xs []int) int {
	result := xs[0]

	for _, n := range xs {
		if n > result {
			result = n
		}
	}

	return result
}
