package parse

import "sort"

func SortTraceInPlace(trace []float64) {
	if len(trace) == 0 {
		return
	}
	sort.Float64s(trace)
}
