package pipelinepattern

import "fmt"

func RunAssemblyLine(numbers ...int) {
	if len(numbers) == 0 {
		numbers = []int{1, 2, 3, 4, 5}
	}

	// Stage 1: Parts Maker
	partsMaker := func(numbers []int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range numbers {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Stage 2: Widget Assembler
	widgetAssembler := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Run the assembly line
	parts := partsMaker(numbers)
	widgets := widgetAssembler(parts)

	// Quality Control
	for widget := range widgets {
		fmt.Printf("Quality Control: Widget %d passed inspection!\n", widget)
	}
}