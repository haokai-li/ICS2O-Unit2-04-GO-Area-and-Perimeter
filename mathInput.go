// Created by: haokai
// Created on: May 2021
// This program displays, "Area-and-Perimeter"


package main

import "fmt"

func main() {

	// This function does addition
	var length int
  var width int
	var area int
	var perimeter int

	// input
	fmt.Println("This program finds the area and perimeter of a rectangle.")
	fmt.Println()
	fmt.Print("Enter the length (cm): ")
	fmt.Scanln(&length)
	fmt.Print("Enter the width (cm): ")
	fmt.Scanln(&width)
	fmt.Println()

	//process
	area = length * width
	perimeter = 2 * (length + width)

	// output
	fmt.Println("The area is: ", area ," cm².")
	fmt.Println("The perimeter is: ", perimeter," cm.")
	fmt.Println("\nDone.")
}