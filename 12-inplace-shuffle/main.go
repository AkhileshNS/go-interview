package main

import "fmt"

func myFunction(arg string) string {

	// Write the body of your function here

	return "running with " + arg
}

func main() {

	// Run your function through some test cases here.
	// Remember: debuggin is half the battle!
	fmt.Println(myFunction("test input"))
}
