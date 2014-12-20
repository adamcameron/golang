// getSubseries.go

package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type SubArrayResult struct {
	Sequence []int
	Total    int
}

var input = flag.String("input", "", "Comma delimited list of integer numbers")
var threshold = flag.Int("threshold", 10, "Integer number for array slice to not exceed")

func main() {
	flag.Parse()

	/*
	 * Array of integers to determine longest sequence from. These are a string
	 * list on the command line that are converted to integers.
	 */
	inputArray := make([]int, 0)

	if *input == "" {
		fmt.Println("Please provide a valid input of comma delimited integers")
		os.Exit(1)
	}

	/*
	 * Get the input from the console args, and split them up
	 * into integers in an array
	 */
	for _, value := range strings.Split(*input, ",") {
		i, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("ERROR - Unable to convert", value, "to integer")
			os.Exit(1)
		}

		inputArray = append(inputArray, i)
	}

	/*
	 * Loop, slice, sum
	 */
	sequences := make([]SubArrayResult, 0)

	for index, _ := range inputArray {
		sequence := make([]int, 0)
		total := 0

		inputSlice := inputArray[index:]

		for _, value := range inputSlice {
			if (total + value) <= *threshold {
				total += value
				sequence = append(sequence, value)
			} else {
				break
			}
		}

		if len(sequence) > 0 {
			result := SubArrayResult{
				sequence,
				total,
			}

			sequences = append(sequences, result)
		}
	}

	/*
	 * Who wins?
	 */
	longestIndex := 0
	longestLength := 0

	for index, item := range sequences {
		if len(item.Sequence) > longestLength {
			longestLength = len(item.Sequence)
			longestIndex = index
		}

		if len(item.Sequence) == longestLength && item.Total > sequences[longestIndex].Total {
			longestLength = len(item.Sequence)
			longestIndex = index
		}
	}

	if len(sequences) > 0 {
		winningSequence := sequences[longestIndex]

		fmt.Println("")
		fmt.Println("Sequence", winningSequence.Sequence, "wins with a length of", len(winningSequence.Sequence), " for a total of", winningSequence.Total)
	} else {
		fmt.Println("No winners :(")
	}
}
