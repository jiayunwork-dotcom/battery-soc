package report

import "battery-soc/internal/parse"

func scanMinMax(trace []float64) (min, max float64) {
	if len(trace) == 0 {
		return 0, 0
	}
	cp := append([]float64(nil), trace...)
	parse.SortTraceInPlace(cp)
	return cp[0], cp[len(cp)-1]
}
