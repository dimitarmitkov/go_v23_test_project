package main

import "fmt"

func Countdown(v int) func(func(int) bool) {
	// next, we return a callback func which is typically
	// called yield, but names like next could also be
	// applicable
	return func(yield func(int) bool) {
		// we then start a for loop that iterates
		for i := v; i >= 0; i-- {
			// once we've finished looping
			if !yield(i) {
				// we then return and finish our iterations
				return
			}
		}
	}
}

func LoopFor() {
	for x := range Countdown(10) {
		fmt.Println(x)
	}

}
