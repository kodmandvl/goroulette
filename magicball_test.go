package goroulette

import (
	"testing"
)

func TestMagicBallAllAnswersAppear(t *testing.T) {
	const trials = 5000

	seen := make(map[string]int, len(magicBallAnswers))

	for i := 0; i < trials; i++ {
		ans := MagicBall("Тестовый вопрос")
		seen[ans]++
	}

	var missing []string
	for _, expected := range magicBallAnswers {
		if _, ok := seen[expected]; !ok {
			missing = append(missing, expected)
		}
	}

	if len(missing) > 0 {
		t.Errorf("за %d попыток не встретились ответы: %v", trials, missing)
	}

	if len(seen) != len(magicBallAnswers) {
		t.Errorf("ожидалось %d уникальных ответов, получено %d: %v",
			len(magicBallAnswers), len(seen), seen)
	}

	// Дополнительно:
	for _, ans := range magicBallAnswers {
		t.Logf("%-45s встретился %d раз(а)", ans, seen[ans])
	}
}
