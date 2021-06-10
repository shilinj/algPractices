package main

import (
	"algstu/interview/algorithm"
	"algstu/learning/routine"
)

func main() {
	// algorithm
	// algorithm.TestBubbleSort()
	// algorithm.TestSelectionSort()
	algorithm.TestQuickSort()

	// interview
	// interview.AlternatelyPrintV2()

	// concurrent
	// concurrence.TestSyncCocurrent()
	// concurrent.TestSyncWaitGroup()
	// concurrent.TestCloseChToBroadcast1()
	routine.TestCloseChToBroadcast2()
}
