package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calSum(param []string, size int) int {
	//len(array) get size of array
	//variable to hold sum
	sum := 0
	count := 0

	//convert to num
	CONVERT: temp, err := strconv.Atoi(param[count])
	//if err in conversion, exit
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	//check to see if the number is positive. if so, add
	if (temp > 0) {
		//square and add to sum
		sum += (temp * temp)
	}

	//if there are still more numbers, go back
	if (count < size - 1) {
		count++
		goto CONVERT
	}

	//return the sum
	return sum
}

func main() {
	// declare variables
	var trials int
	count := 0 //keep track of the number of times
	scanner := bufio.NewScanner(os.Stdin)

	//get number of trials
	fmt.Scan(&trials)

	//make arrays to store the results and number of numbers for each trial
	num_in_trials := make([]int, trials)
	results := make([]int, trials)

	//This is the spot to go back to to read next line
	SCAN: scanner.Scan()
	//convert number of numbers from string to number and store in array
	temp, err := strconv.Atoi(scanner.Text())
	//if err in conversion, exit
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	num_in_trials[count] = temp
	//read the line with the numbers
	scanner.Scan()

	//split
	working := strings.Split(scanner.Text(), " ")

	//send to the function to calculate sum
	results[count] = calSum(working, num_in_trials[count])

	//if there are more trials return to scan next two lines
	if (count < trials - 1) {
		count++
		goto SCAN
	}

	//reset the count variable
	count = 0
	//Start printing the results now that all the inputs have been read
	PRINT: fmt.Println(results[count])

	//if there is still more to print, go back
	if (count < trials - 1) {
		count++
		goto PRINT
	}
}