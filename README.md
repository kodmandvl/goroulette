# goroulette

Golang demo package with roulette game and magic ball

My similar extension for PostgreSQL: [roulette_extension](https://github.com/kodmandvl/roulette_extension)

Add new tag (after commit and push):

```bash
git tag v0.0.7
git push origin v0.0.7
```

## Example

```golang
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kodmandvl/goroulette"
)

func main() {
	var max, bet, question string
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: goroulette roulette|magicball")
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "roulette":
		fmt.Println("===== ROULETTE =====")
		fmt.Print("Maximum value: ")
		fmt.Scan(&max)
		m, errm := strconv.Atoi(max)
		if errm != nil {
			fmt.Println("Maximum value must be integer!")
			return
		}
		for {
			fmt.Print("Your bet: ")
			fmt.Scan(&bet)
			b, errb := strconv.Atoi(bet)
			if errb != nil {
				fmt.Println("Your bet must be integer!")
				return
			}
			fmt.Println(goroulette.Roulette(m, b))
		}
	case "magicball":
		fmt.Println("===== MAGIC BALL =====")
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("Your question: ")
			// Читаем всё до нажатия Enter ('\n')
			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Input error:", err)
				return
			}
			// Убираем символы переноса строки (\n в Linux/macOS и \r\n в Windows)
			question = strings.TrimRight(input, "\r\n")
			// fmt.Printf("Your question: [%s]\n", question)
			fmt.Printf("Answer: %s\n", goroulette.MagicBall(question))
		}
	default:
		fmt.Println("Not a valid option.")
		fmt.Println("Usage: goroulette roulette|magicball")
	}
}
```