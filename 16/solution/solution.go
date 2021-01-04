package solution

import (
	"strings"

	"github.com/t3h2mas/advent-2020/16/tickets"
)

func TicketScanningErrorRate(input string) int {
	parts := strings.Split(input, "\n\n")

	rules := tickets.ParseRules(parts[0])
	tickets := tickets.ParseTickets(parts[2])
	errorRate := 0

	for _, ticket := range tickets {
		for _, value := range ticket.Values() {
			validAtLeastOnce := false
			for _, rule := range rules {
				if rule.Valid(value) {
					validAtLeastOnce = true
					break
				}
			}

			if !validAtLeastOnce {
				errorRate += value
			}
		}
	}

	return errorRate
}
