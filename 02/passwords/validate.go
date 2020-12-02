package passwords

func ValidPasswordsByPolicy(entries []string, validator Validator) int {
	validCount := 0

	for _, entry := range entries {
		corporatePassword := corporatePasswordFrom(entry)

		corporatePassword.WithValidator(validator)

		if corporatePassword.IsValid() {
			validCount++
		}
	}

	return validCount
}
