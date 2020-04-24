package main

import "fmt"

const (
	PassingColor  = "\033[1;34mPASSING\033[0m"
	ErrorColor    = "\033[1;31mFAILED\033[0m"
	WellDoneColor = "\033[1;34mWELL DONE!\033[0m"
)

func runTaskTest(fn func(interface{}) interface{}, testInputs []interface{}, testOutputs []interface{}, taskName string) (int, int) {
	fmt.Println("======= ", taskName, "\t Running tests =======")
	var passing, failed int = 0, 0
	for index, s := range testInputs {
		var result = fn(s)
		fmt.Print(s, "\t=>\t", result, "\t")
		if testOutputs[index] == result {
			fmt.Println(PassingColor)
			passing++
		} else {
			fmt.Println(ErrorColor, "expected: ", testOutputs[index])
			failed++
		}
	}
	fmt.Println("Tests passed: ", passing, "/", failed)
	fmt.Print("Result: ")
	if failed == 0 {
		fmt.Println(WellDoneColor)
	} else {
		fmt.Println(ErrorColor)
	}
	return passing, failed
}

// wrapping function that returns a string
// in one that returns empty interface
// so as to sidestep Go's lack of generic types
func task1runner(s interface{}) interface{} {
	str, ok := s.(string)
	if !ok {
		return false
	}
	return palindrome(str)
}

func task4runner(n interface{}) interface{} {
	nn, ok := n.([]int)
	if !ok {
		return false
	}
	return fib(nn)
}

func performTests() {
	var passing, failed int = 0, 0

	//task 1
	testStrings := []interface{}{"wow", "much", "pallap", "kayak", "oo"}
	expectedValues := []interface{}{true, false, true, true, true}
	var p, f int = runTaskTest(task1runner, testStrings, expectedValues, "1. Palindrome")
	passing += p
	failed += f

	//task 2
	//TODO
	testInts := []interface{}{[]int{1, 2, 3}, []int{2, 3, 8}}
	expectedInts := []interface{}{4, 24}
	p, f = runTaskTest(task4runner, testInts, expectedInts, "4. Fib")
	passing += p
	failed += f

	fmt.Println("======= \t RESULTS SUMMARY \t======")
	fmt.Println(PassingColor, ":\t", passing)
	fmt.Println(ErrorColor, ":\t", failed)
}
