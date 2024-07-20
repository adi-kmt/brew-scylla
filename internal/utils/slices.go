package utils

import (
	"reflect"
	"strconv"
)

func Contains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func GetFieldSliceFromEntitySlice[T any](s []T, field string) []string {
	var result []string

	for _, entity := range s {
		fieldValue := reflect.ValueOf(entity).FieldByName(field)

		if fieldValue.IsValid() {
			var valueString string
			switch fieldValue.Kind() {
			case reflect.String:
				valueString = fieldValue.String()
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				valueString = strconv.FormatInt(fieldValue.Int(), 10)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				valueString = strconv.FormatUint(fieldValue.Uint(), 10)
			case reflect.Float32, reflect.Float64:
				valueString = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
			default:
				continue
			}
			result = append(result, valueString)
		} else {
			continue
		}
	}

	return result
}
