package goroulette

import (
	"fmt"
	"strings"
	"testing"
)

func TestRouletteAllNumbersCanWin(t *testing.T) {
	const max = 36
	const trials = 5000

	type stat struct {
		bet  int
		wins int
	}

	var results []stat
	var neverWon []int

	for bet := 1; bet <= max; bet++ {
		wins := 0

		for i := 0; i < trials; i++ {
			res := Roulette(max, bet)
			if strings.Contains(res, "VICTORY") {
				wins++
			}
		}

		results = append(results, stat{bet: bet, wins: wins})

		if wins == 0 {
			neverWon = append(neverWon, bet)
		}
	}

	if len(neverWon) > 0 {
		t.Errorf("следующие числа не выиграли ни разу за %d попыток: %v",
			trials, neverWon)
	}

	// Дополнительно:
	for _, r := range results {
		t.Logf("число %2d: выиграло %4d раз из %d",
			r.bet, r.wins, trials)
	}
}

func TestRouletteOutcomeFormat(t *testing.T) {
	const max = 36
	const bet = 7
	const trials = 5000

	gotWin := false
	gotLoss := false

	wantWin := fmt.Sprintf("YOU GET YOUR NUMBER %d!", bet)
	wantLoss := fmt.Sprintf("your number %d did not come up", bet)

	for i := 0; i < trials; i++ {
		res := Roulette(max, bet)

		switch {
		case strings.Contains(res, wantWin):
			gotWin = true
		case strings.Contains(res, wantLoss):
			gotLoss = true
		default:
			t.Fatalf("неожиданный формат ответа: %q", res)
		}

		if gotWin && gotLoss {
			break
		}
	}

	if !gotWin {
		t.Error("за 5000 попыток не встретилась ни одна победа (крайне подозрительно)")
	}
	if !gotLoss {
		t.Error("за 5000 попыток не встретилось ни одного проигрыша (невозможная ситуация)")
	}
}
