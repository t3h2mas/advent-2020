package customs

type FormAnswers map[rune]bool

func NewFormAnswers() FormAnswers {
	return make(map[rune]bool)
}

func (f FormAnswers) MarkYes(question rune) {
	f[question] = true
}

func ParseGroupAnswers(answers []string) FormAnswers {
	result := NewFormAnswers()

	for _, answer := range answers {
		for _, question := range answer {
			result.MarkYes(question)
		}
	}

	return result
}
