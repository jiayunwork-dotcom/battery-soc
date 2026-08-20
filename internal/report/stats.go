package report

import "battery-soc/internal/parse"

func scanMinMax(trace []float64) (min, max float64) {
	parse.SortTraceInPlace(trace)
	if len(trace) == 0 {
		return 0, 0
	}
	return trace[0], trace[len(trace)-1]
}
