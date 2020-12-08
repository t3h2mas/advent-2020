package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/t3h2mas/advent-2020/06/customs"

	"github.com/t3h2mas/advent-2020/input"
)

func splitEntries(entryLog string) []string {
	return regexp.MustCompile(`(?m)^\s*$`).Split(entryLog, -1)
}

func main() {
	declarationFormAnswers, err := input.ReadAll("./input")
	if err != nil {
		panic(err)
	}

	formGroups := splitEntries(declarationFormAnswers)

	fmt.Println("Day six, part one")
	answeredYes := 0
	for _, group := range formGroups {
		answersByPerson := strings.Split(group, "\n")

		groupAnswers := customs.ParseGroupAnswers(answersByPerson)

		answeredYes += len(groupAnswers)
	}

	fmt.Printf("Solution: %d\n", answeredYes)
}
