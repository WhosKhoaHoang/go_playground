package main

import (
	"fmt"
	"go_playground/utils"
)

func main() {
	test := utils.Dude{X:"a", Y:"b"}
	fmt.Println("SNAPPPPPP")
	fmt.Println(test)

	// test123 := utils.TestReadFile("test.json")
	// fmt.Println("DOES IT WORK?!")
	// fmt.Println(test123)

	// target123 := utils.TestReadTarget("target.json")
	// fmt.Println("DOES TARGET STUFF WORK?!")
	// fmt.Println(target123)

	utils.TestReadUser("users.json")
}