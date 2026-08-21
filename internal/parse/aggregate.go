package parse

import "battery-soc/internal/coulomb"

func addCharge(samp coulomb.CurrentSample) float64 {
	return samp.Current * samp.DT
}
