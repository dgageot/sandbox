package main

func CountCorrectColorAndPosition(secret [4]string, guess [4]string) int {
	count := 0

	for i := range secret {
		if secret[i] == guess[i] {
			count++
		}
	}

	return count
}

func CountCorrectColorWrongPosition(secret [4]string, guess [4]string) int {
	count := 0

	colors := allColors(secret, guess)
	for _, color := range colors {
		count += countWrongPosition(color, secret, guess)
	}

	return count
}

func allColors(secret [4]string, guess [4]string) []string {
	colors := make(map[string]bool)

	for _, color := range secret {
		colors[color] = true
	}
	for _, color := range guess {
		colors[color] = true
	}

	allColors := make([]string, len(colors))
	for color := range colors {
		allColors = append(allColors, color)
	}
	return allColors
}

func countWrongPosition(color string, secret [4]string, guess [4]string) int {
	inSecret := 0
	inGuess := 0

	for i := range secret {
		if secret[i] != guess[i] {
			if secret[i] == color {
				inSecret++
			} else if guess[i] == color {
				inGuess++
			}
		}
	}

	if inSecret < inGuess {
		return inSecret
	}
	return inGuess
}
