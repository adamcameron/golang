// Subseries.go

package subseries

func Get(series []int, threshold int) []int {
	subseries := series[0:0]
	low := 0
	for index, _ := range series {
		high := index + 1
		working := series[low:high]

		for sumSlice(&working) > threshold {
			low++
			working = series[low:high]
		}

		if len(working) > len(subseries) {
			subseries = working
		} else if (len(working) == len(subseries)) && (sumSlice(&working) > sumSlice(&subseries)) {
			subseries = working
		}
	}
	return subseries
}

func sumSlice(slice *[]int) (sum int) {
	for _, value := range *slice {
		sum += value
	}
	return
}
