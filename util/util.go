package util

import (
	"encoding/base64"
	"reflect"
)

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

		v := reflect.ValueOf(param)
		switch v.Kind() {
		case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
			if v.IsNil() {
				return true
			}
		}
	}
	return false
}

func StringToBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
