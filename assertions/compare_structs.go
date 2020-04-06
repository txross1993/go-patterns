package assertions

import (
	"reflect"
	"testing"
)

// CompareOnlyNonNilTestFields satisfies assert's ComparisonAssertionFunc type
// and will compare fields between two structs
func CompareOnlyNonNilTestFields(t *testing.T, expected interface{}, actual interface{}) bool {
	elements := reflect.ValueOf(&expected).Elem() // type reflect.StructField
	typeOfExpected := elements.Type()

	expectedValue := reflect.ValueOf(expected) // type reflect.Value
	actualValue := reflect.ValueOf(actual)     // type reflect.Value

	for i := 0; i < elements.NumField(); i++ {
		fieldName := typeOfExpected.Field(i).Name
		if actualValue.FieldByName(fieldName) != expectedValue.FieldByName(fieldName) {
			return false
		}
	}

	return true
}
