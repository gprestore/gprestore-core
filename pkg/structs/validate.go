package structs

import (
	"reflect"
)

func IsEmpty(input any) bool {
	structValue := reflect.Indirect(reflect.ValueOf(input))

	var results []bool
	numField := structValue.NumField()
	for i := range numField {
		val := structValue.Field(i)

		isEmpty := val.IsValid() && val.IsZero()
		results = append(results, isEmpty)
	}

	numEmpty := 0
	for _, r := range results {
		if r {
			numEmpty += 1
		}
	}

	return numField == numEmpty
}
