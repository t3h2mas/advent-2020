package passwords

import (
	"strconv"
	"strings"
)

type Validator interface {
	Validate(corporatePassword *CorporatePassword) bool
}

type PasswordPolicy struct {
	min  int
	max  int
	char rune
}

type CorporatePassword struct {
	password  string
	policy    *PasswordPolicy
	validator Validator
}

func (cp *CorporatePassword) WithValidator(validator Validator) {
	cp.validator = validator
}

func (cp *CorporatePassword) IsValid() bool {
	return cp.validator.Validate(cp)
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
