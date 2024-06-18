package util

func BoolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func AnyNil(params ...interface{}) bool {
	for _, param := range params {
		if param == nil {
			return true
		}
	}
	return false
}
