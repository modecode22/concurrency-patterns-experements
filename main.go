package main

import (
	"fmt"
	"github.com/modecode22/concurrency-patterns-experements/donepattern"
	"github.com/modecode22/concurrency-patterns-experements/forselectpattern"
	"github.com/modecode22/concurrency-patterns-experements/pipelinepattern"
)

func main() {
	fmt.Println("Welcome to the Go Concurrency Circus! 🎪")

	fmt.Println("\n🤹 For-Select Juggling Act:")
	forselectpattern.PerformJugglingAct()

	fmt.Println("\n🚪 Done Channel Magic Door:")
	donepattern.PerformMagicDoorTrick()

	fmt.Println("\n🏭 Pipeline Assembly Line:")
	pipelinepattern.RunAssemblyLine()

	fmt.Println("\nThat's all, folks! Remember, concurrency is like cooking - timing is everything!")
}