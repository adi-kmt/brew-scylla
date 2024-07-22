package utils

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

func SliceContains[T comparable](s []T, item T) bool {
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

func GetEntityThatMatchesInSlice[T any](s []T, field string, value string) (*T, error) {
	for _, entity := range s {
		fieldValue := reflect.ValueOf(entity).FieldByName(field)
		if fieldValue.IsValid() {
			if fieldValue.String() == value {
				return &entity, nil
			}
		}
	}
	return nil, fmt.Errorf("entity not found")
}

func SearchEntityFieldFromSlice[T any](s []T, field string, pattern string) ([]T, error) {
	listOfStrings := GetFieldSliceFromEntitySlice(s, field)

	threshold := 2

	var matchedItems []string

	for _, str := range listOfStrings {
		dist := levenshtein.DistanceForStrings([]rune(pattern), []rune(str), levenshtein.DefaultOptions)
		if dist <= threshold {
			matchedItems = append(matchedItems, str)
		}
	}

	if len(matchedItems) > 0 {
		var result []T
		for _, item := range matchedItems {
			entity, err := GetEntityThatMatchesInSlice(s, field, item)
			if err != nil {
				return nil, err
			}
			result = append(result, *entity)
		}
		return result, nil
	}
	return nil, nil
}
