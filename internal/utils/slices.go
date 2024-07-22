package utils

import (
	"fmt"
	"reflect"
	"strconv"
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

// Based on Bitap Algorithm or Baeza-Yates-Gonnet Algorithm for fuzzy string matching.
func GetFuzzyMatchedStringSliceFromSlice(texts []string, pattern string, k int) []string {
	var matches []string

	if pattern == "" {
		return texts // Return all texts if pattern is empty
	}
	if len(pattern) > 31 {
		return []string{"The pattern is too long!"} // Handle excessively long patterns
	}

	for _, text := range texts {
		if bitapFuzzyBitwiseSearch(text, pattern, k) {
			matches = append(matches, text)
		}
	}

	return matches
}

func bitapFuzzyBitwiseSearch(text string, pattern string, k int) bool {
	m := len(pattern)

	if m == 0 {
		return true // An empty pattern matches any text
	}
	if m > 31 {
		return false // Pattern length exceeds the limit
	}

	// Initialize the bit arrays
	var R []uint64
	patternMask := make([]uint64, 256)

	R = make([]uint64, k+1)
	for i := 0; i <= k; i++ {
		R[i] = ^uint64(1)
	}

	for i := 0; i < 256; i++ {
		patternMask[i] = ^uint64(0)
	}
	for i := 0; i < m; i++ {
		patternMask[pattern[i]] &= ^(uint64(1) << uint64(i))
	}

	// Bitap algorithm implementation
	for i := 0; i < len(text); i++ {
		oldRd1 := R[0]

		R[0] |= patternMask[text[i]]
		R[0] <<= 1

		for d := 1; d <= k; d++ {
			tmp := R[d]
			R[d] = (oldRd1 & (R[d] | patternMask[text[i]])) << 1
			oldRd1 = tmp
		}

		if 0 == (R[k] & (uint64(1) << uint64(m))) {
			return true // Match found
		}
	}

	return false // No match found
}
