package cardWithCopies

import (
	cards "gitlab.com/alexmarkelov42/AdventOfCode/2023/day_4/part1"
)

func GetTotalPointsWithCopies(winNumbers, cardNumbers [][]int) int {
	var (
		matches = 0
		result = 0
	)
	cardCopies := make([]int, len(cardNumbers))
	for i := range cardCopies {
		cardCopies[i] = 1
	}
	for i := range cardCopies {
		matches = cards.GetNumberOfMatches(winNumbers[i], cardNumbers[i])
		if matches > 0 {
			cardCopies = addNewCardCopies(cardCopies, i, matches, len(winNumbers))
		}
	}
	for _, num := range cardCopies {
		result += num
	}
	return result
}

func addNewCardCopies(cardCopies []int, i, matches, lenNumbers int) []int {
	for k := i + 1; k <= matches+i; k++ {
		if k == lenNumbers {
			break
		} else {
			cardCopies[k]+=cardCopies[i]
		}

	}
	return cardCopies
}
