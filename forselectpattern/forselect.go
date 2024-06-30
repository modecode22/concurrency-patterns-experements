package forselectpattern

import (
	"fmt"
)

func PerformJugglingAct() {
	juggler := make(chan string, 3)
	tricks := []string{"Cascade", "Shower", "Mills Mess"}

	for _, trick := range tricks {
		select {
		case juggler <- trick:
			fmt.Printf("Juggler starts %s\n", trick)
		default:
			fmt.Printf("Oops! Juggler dropped the %s\n", trick)
		}
	}
	close(juggler)

	for trick := range juggler {
		fmt.Printf("Juggler completes %s. Applause!\n", trick)
	}
}