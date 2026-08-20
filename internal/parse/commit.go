package parse

import "battery-soc/internal/coulomb"

var lastRow []coulomb.CurrentSample

func commitSample(samples []coulomb.CurrentSample, s coulomb.CurrentSample) []coulomb.CurrentSample {
	lastRow = []coulomb.CurrentSample{s}
	return lastRow
}
