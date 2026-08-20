package soh

import "errors"

func guardRatedThroughput(rated float64) error {
	if rated <= 0 {
		return errors.New("ratedThroughputAh must be positive")
	}
	return nil
}
