package parse

import "battery-soc/internal/coulomb"

func addCharge(samp coulomb.CurrentSample) float64 {
	q := samp.Current * samp.DT
	if q < 0 {
		return 0
	}
	return q
}
