package passwords

import (
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	min  int
	max  int
	char rune
}

type CorporatePassword struct {
	password string
	policy   *PasswordPolicy
}

func (cp *CorporatePassword) IsValid() bool {
	charCount := 0
	for _, ch := range cp.password {
		if ch == cp.policy.char {
			charCount++
		}

		if charCount > cp.policy.max {
			return false
		}
	}

	return charCount >= cp.policy.min
}

// parse 1-3 b: cdefg into corporate password
func corporatePasswordFrom(entry string) *CorporatePassword {
	parts := strings.Split(entry, ": ")
	password := parts[1]

	minMaxChar := strings.Split(parts[0], " ")

	char := minMaxChar[1][0]

	minMax := strings.Split(minMaxChar[0], "-")

	min, err := strconv.Atoi(minMax[0])

	if err != nil {
		panic(err)
	}

	max, err := strconv.Atoi(minMax[1])

	if err != nil {
		panic(err)
	}

	return &CorporatePassword{
		password: password,
		policy: &PasswordPolicy{
			min:  min,
			max:  max,
			char: rune(char),
		},
	}
}

func ValidPasswordsByPolicy(entries []string) int {
	validCount := 0

	for _, entry := range entries {
		corporatePassword := corporatePasswordFrom(entry)

		if corporatePassword.IsValid() {
			validCount++
		}
	}

	return validCount
}
