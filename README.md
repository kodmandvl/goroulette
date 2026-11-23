# goroulette

Golang demo package with roulette game and magic ball

My similar extension for PostgreSQL: [roulette_extension](https://github.com/kodmandvl/roulette_extension)

Add new tag (after commit and push):

```bash
git tag v0.0.4
git push origin v0.0.4
```

## Example

```golang
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/kodmandvl/goroulette"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: roulette|magicball <arguments>")
		return
	}

	// Differentiating between the commands
	switch arguments[1] {
	case "roulette":
		if len(arguments) < 4 {
			fmt.Println("Usage: roulette <max> <bet>")
			return
		}
		m := arguments[2]
		b := arguments[3]
		max, okm := strconv.Atoi(m)
		bet, okb := strconv.Atoi(b)
		if okm != nil || okb != nil {
			fmt.Println("Usage: roulette <max> <bet>")
			return
		}
		fmt.Println("Max value:", max)
		fmt.Println("Your bet:", bet)
		fmt.Println(goroulette.Roulette(max, bet))
	case "magicball":
		if len(arguments) < 3 {
			fmt.Println("Usage: magicball 'Your question?'")
			return
		}
		q := arguments[2]
		fmt.Println("Your question:", q)
		fmt.Printf("Answer: %s\n", goroulette.MagicBall(q))
	default:
		fmt.Println("Not a valid option")
	}
}
```