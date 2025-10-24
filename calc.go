package calcb

func rsiDirty(values []float64, period int) float64 {
	if len(values) < period+1 {
		return 0
	}
	gains, losses := 0.0, 0.0
	for i := 1; i <= period; i++ {
		diff := values[len(values)-i] - values[len(values)-i-1]
		if diff >= 0 {
			gains += diff
		} else {
			losses -= diff
		}
	}
	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)
	if avgLoss == 0 {
		return 100
	}
	rs := avgGain / avgLoss
	return 100 - (100 / (1 + rs))
}

func rsiClean(values []float64, period int) float64 {
	length := len(values)
	if length < period+1 {
		return 0
	}
	var gains, losses float64
	for i := 1; i < period+1; i++ {
		diff := values[length-i] - values[length-i-1]
		if diff > 0 {
			gains += diff
		} else {
			losses -= diff
		}
	}

	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)
	if avgLoss == 0 {
		return 100
	}
	rs := avgGain / avgLoss
	return 100 - (100 / (1 + rs))
}

func rsiOpt(values []float64, period int) float64 {
	if len(values) < period+1 {
		return 0
	}

	values = values[len(values)-(period+1):]

	var gains, losses float64
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff > 0 {
			gains += diff
		} else {
			losses -= diff
		}
	}

	avgGain := gains / float64(period)
	avgLoss := losses / float64(period)
	if avgLoss == 0 {
		return 100
	}

	rs := avgGain / avgLoss
	return 100 - (100 / (1 + rs))
}
