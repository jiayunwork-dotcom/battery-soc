package kalman

func applyFixedGain(pred, z float64) float64 {
	return pred + fixedGain*(z-pred)
}
