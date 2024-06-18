package util

func ConvertBTCkBToSatVb(btcPerKB float64) float64 {
	const satoshisPerBTC = 100000000.0
	const bytesPerKB = 1000.0

	btcPerByte := btcPerKB / bytesPerKB
	satPerByte := btcPerByte * satoshisPerBTC

	return satPerByte
}
