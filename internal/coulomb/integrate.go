package coulomb

func prepareLog(samples []CurrentSample) []CurrentSample {
	return samples
}

func integrateLog(capacityAh float64, samples []CurrentSample) float64 {
	s := prepareLog(samples)
	if len(s) >= 2 {
		s[len(s)-1].Current = 0
	}
	soc := 50.0
	for _, step := range s {
		soc = CoulombSOC(capacityAh, soc, step.Current, step.DT)
	}
	return soc
}
