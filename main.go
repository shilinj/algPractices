package main

import (
	"algstu/interview"
	"algstu/interview/algorithm"
	"algstu/learning"
	"algstu/learning/routine"
)

func main() {
	// algorithm
	algorithm.TestBubbleSort()
	algorithm.TestSelectionSort()
	algorithm.TestQuickSort()

	// interview
	interview.AlternatelyPrintV1()

	// concurrent
	routine.TestSyncCocurrent()
	routine.TestSyncWaitGroup()
	routine.TestCloseChToBroadcast1()
	routine.TestCloseChToBroadcast2()

	// learning
	learning.TestDefer()
	learning.TestSlice()
	learning.TestArgs()
}
