// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-25
// Fileoverview: This program calculates the average of a set of marks
// and displays a performance message.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// variable declarations
	var numberOfMarks int
	var markInput float64
	var sum float64
	var average float64

	reader := bufio.NewReader(os.Stdin)

	// ask for number of marks
	fmt.Print("How many marks will you enter today? ")
	numberString, _ := reader.ReadString('\n')
	numberString = strings.TrimSpace(numberString)
	numberOfMarks, _ = strconv.Atoi(numberString)

	// loop to get all marks
	for counter := 1; counter <= numberOfMarks; counter++ {
		fmt.Printf("Enter mark %d: ", counter)
		markString, _ := reader.ReadString('\n')
		markString = strings.TrimSpace(markString)
		markInput, _ = strconv.ParseFloat(markString, 64)

		sum = sum + markInput
	}

	// calculate average
	average = sum / float64(numberOfMarks)

	// display results
	fmt.Printf("You have entered %d marks. The student's average is %d%%.\n",
		numberOfMarks, int(average+0.5))

	// performance message
	if average <= 49 {
		fmt.Println("The student is failing.")
	} else if average >= 50 && average <= 69 {
		fmt.Println("The student's performance is just under average.")
	} else if average >= 70 && average <= 79 {
		fmt.Println("The student's performance is average.")
	} else {
		fmt.Println("The student is on the honour roll.")
	}

	fmt.Println("\nDone.")
}
