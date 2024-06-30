package donepattern

import (
	"fmt"
	"time"
)

func PerformMagicDoorTrick() {
	magicDoor := make(chan bool)

	go func() {
		for {
			select {
			case <-magicDoor:
				fmt.Println("Magician: Abracadabra! The door vanishes!")
				return
			default:
				fmt.Println("Magician: The door is still here...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	close(magicDoor)
	fmt.Println("Audience: Wow! The door disappeared!")
}
