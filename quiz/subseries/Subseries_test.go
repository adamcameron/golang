// Subseries_test.go

package subseries

import (
	"reflect"
	"testing"
)

func TestReturnsElementIfWithinThreshold(t *testing.T) {
	series := []int{100}
	threshold := 100
	expected := []int{threshold}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestReturnsMultipleElementIfWithinThreshold(t *testing.T) {
	series := []int{100, 100}
	threshold := 500
	expected := []int{100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestTotalElementsWithinThreshold(t *testing.T) {
	series := []int{100, 100, 100, 100, 100, 100}
	threshold := 500
	expected := []int{100, 100, 100, 100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)

	total := sumSlice(&result)
	if total > threshold {
		t.Errorf("\ntotal %d should be less than threshold %d", total, threshold)
	}

}

func TestReturnsSubsequentLargerSubseries(t *testing.T) {
	series := []int{600, 100, 100, 100, 600, 100, 100, 100, 100, 600}
	threshold := 500
	expected := []int{100, 100, 100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestReturnsLongerAdjacentSubseries(t *testing.T) {
	series := []int{600, 100, 100, 100, 600, 200, 100, 100, 100, 100, 100, 600}
	threshold := 500
	expected := []int{100, 100, 100, 100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWithEmptySeries(t *testing.T) {
	series := []int{}
	threshold := 500
	expected := []int{}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWhenThresholdIsLessThanAllItems(t *testing.T) {
	series := []int{600, 700, 800, 900}
	threshold := 500
	expected := []int{}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWhenThresholdIsGreaterThanAllItems(t *testing.T) {
	series := []int{600, 700, 800, 900}
	threshold := 1000
	expected := []int{900}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWhenThresholdIsGreaterThanSumOfAllItems(t *testing.T) {
	series := []int{600, 700, 800, 900}
	threshold := 5000
	expected := []int{600, 700, 800, 900}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWhenTheSubseriesIsEqualToTheThresholdAndIsAtTheBeginningOfTheSeries(t *testing.T) {
	series := []int{100, 100, 100, 100, 100, 100, 600, 150, 150, 150, 150}
	threshold := 600
	expected := []int{100, 100, 100, 100, 100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksWhenTheSubseriesIsEqualToTheThresholdAndIsAtTheEndOfTheSeries(t *testing.T) {
	series := []int{600, 150, 150, 150, 150, 100, 100, 100, 100, 100, 100}
	threshold := 600
	expected := []int{100, 100, 100, 100, 100, 100}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestWorksAsPerStatedRequirement(t *testing.T) {
	series := []int{100, 300, 100, 50, 50, 50, 50, 50, 500, 200, 100}
	threshold := 500
	expected := []int{100, 50, 50, 50, 50, 50}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestReturnsSubseriesWithHighestTallyWhenThereAreTwoEqualLengthSubseriesAndSecondSubseriesIsHigher(t *testing.T) {
	series := []int{100, 50, 50, 50, 50, 50, 500, 100, 60, 60, 60, 60, 60, 500}
	threshold := 500
	expected := []int{100, 60, 60, 60, 60, 60}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func TestReturnsSubseriesWithHighestTallyWhenThereAreTwoEqualLengthSubseriesAndFirstSubseriesIsHigher(t *testing.T) {
	series := []int{100, 60, 60, 60, 60, 60, 500, 100, 50, 50, 50, 50, 50, 500}
	threshold := 500
	expected := []int{100, 60, 60, 60, 60, 60}
	result := Get(series, threshold)

	assertEquals(t, expected, result)
}

func assertEquals(t *testing.T, expected []int, actual []int) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nexpected %d\n actual %d", expected, actual)
	}
}
