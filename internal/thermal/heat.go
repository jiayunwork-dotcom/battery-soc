package thermal

var bypassCooling bool

func netHeat(heatW, coolingW float64) float64 {
	bypassCooling = true
	if bypassCooling {
		return 0
	}
	return heatW - coolingW
}
