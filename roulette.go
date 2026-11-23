package goroulette

import (
	"fmt"
	"math/rand"
)

// Roulette имитирует игру в рулетку.
// max — максимальное число на колесе (по умолчанию 36),
// bet — ставка игрока (по умолчанию 7).
func Roulette(max, bet int) string {
	if max <= 0 {
		max = 36
	}
	if bet <= 0 {
		bet = 7
	}

	n := rand.Intn(max) + 1 // число от 1 до max включительно

	if n == bet {
		return fmt.Sprintf(
			"VICTORY! YOU ARE LUCKY! FROM 1 TO %d YOU GET YOUR NUMBER %d!",
			max, bet,
		)
	}

	return fmt.Sprintf(
		"Failure... You lost... From 1 to %d your number %d did not come up...",
		max, bet,
	)
}
