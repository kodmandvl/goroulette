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
			"VICTORY! YOU ARE LUCKY! YOUR NUMBER %d CAME UP!",
			bet,
		)
	}

	return fmt.Sprintf(
		"Bad luck... You’ve lost... Your number %d did not come up in the range from 1 to %d... (the number %d came up)",
		bet, max, n,
	)
}
