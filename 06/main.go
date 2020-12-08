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
	// For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?
	answeredYes := 0
	for _, group := range formGroups {
		answersByPerson := strings.Split(group, "\n")

		groupAnswers := customs.ParseGroupAnswers(answersByPerson)

		answeredYes += len(groupAnswers)
	}

	fmt.Printf("Solution: %d\n", answeredYes)

	// For each group, count the number of questions to which everyone answered "yes". What is the sum of those counts?
	fmt.Println("Day six, part two")
	allAnsweredYes := 0
	for _, group := range formGroups {
		fmt.Printf("Starting group '%s'\n", strings.TrimSpace(group))
		answersByPerson := strings.Split(strings.TrimSpace(group), "\n")
		fmt.Printf("%+v\n", answersByPerson)

		var groupAnswerIntersection customs.FormAnswers

		for _, answers := range answersByPerson {
			formAnswers := customs.ParseIndividualAnswers(answers)
			fmt.Printf("From: '%s'\n", answers)
			fmt.Printf("Processing answers... %+v\n", formAnswers)

			if groupAnswerIntersection == nil {
				groupAnswerIntersection = formAnswers
				fmt.Printf("groupAnswerIntersection initialized: %+v\n", groupAnswerIntersection)
				continue
			}

			// apply intersection of current form answers to group answers
			for question := range groupAnswerIntersection {
				if !formAnswers.Has(question) {
					fmt.Printf("Removing '%c' from group\n", question)
					// an individual did not answer yes to question, remove from group
					delete(groupAnswerIntersection, question)
				}
			}
		}

		fmt.Printf("len of groupAnswerIntersection: %d\n", len(groupAnswerIntersection))

		allAnsweredYes += len(groupAnswerIntersection)
	}

	fmt.Printf("Solution: %d\n", allAnsweredYes)
}
