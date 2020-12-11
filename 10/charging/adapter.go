package charging

import (
	"fmt"
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
	sort.Ints(adapters)
	// add the outlet device voltage of 0
	adapters = append([]int{0}, adapters...)

	// add the built in adapter voltage of 3 + the max voltage
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	available := make([]int, adapters[len(adapters)-1]+1)

	for _, voltage := range adapters {
		available[voltage] = 1
	}

	dp := make([]int, len(available))
	dp[0] = 1
	dp[1] = 1

	for i := 2; i < len(available); i++ {
		for j := 1; j < 4; j++ {
			if i-j < 0 {
				continue
			}

			if available[i-1] == 0 {
				dp[i] = 0
			} else {
				dp[i] += dp[(i - j)]
			}
		}
	}

	fmt.Printf("%+v\n", adapters)
	fmt.Printf("%+v\n", available)
	fmt.Printf("%+v\n", dp)
	return dp[len(available)-1]
}
