// getSubseries.go

package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var seriesAsString string
var threshold int

func init() {
	flag.StringVar(&seriesAsString, "series", "", "Series from which to extract subseries")
	flag.IntVar(&threshold, "threshold", 0, "Threshold within which the subseries total must be")
	flag.Parse()
}

func main() {
	flag.Parse()
	series := convertSeriesToArray(seriesAsString)

	subseries := series[0:0]
	low := 0

	for index, _ := range series {
		high := index + 1
		working := series[low:high]

		for sumSlice(&working) > threshold {
			low++
			working = series[low:high]
		}
		sum := sumSlice(&working)

		workingIsBetterForLength := len(working) > len(subseries)
		workingIsBetterForTotal := len(working) == len(subseries) && sum <= threshold

		if workingIsBetterForLength || workingIsBetterForTotal {
			subseries = working
		}
	}

	fmt.Println(subseries)
}

func convertSeriesToArray(seriesAsString string) (seriesAsIntegers []int) {
	for _, stringElement := range strings.Split(seriesAsString, ",") {
		integerElement, e := strconv.Atoi(stringElement)
		if e == nil {
			seriesAsIntegers = append(seriesAsIntegers, integerElement)
		}
	}
	return
}

func sumSlice(slice *[]int) (sum int) {
	for _, value := range *slice {
		sum += value
	}
	return
}
