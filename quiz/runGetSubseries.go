// runGetSubseries.go

package main

import (
	"flag"
	"fmt"
	"github.com/daccfml/golang/quiz/subseries"
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

	subseries := subseries.Get(series, threshold)

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
