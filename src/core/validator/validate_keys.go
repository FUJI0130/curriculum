package validator

import (
	"fmt"
	"reflect"

	"github.com/FUJI0130/curriculum/src/core/validator/customerrors"
)

func ValidateKeysAgainstStruct(rawReq map[string]interface{}, referenceStruct interface{}) error {
	expectedKeys := make(map[string]bool)

	val := reflect.ValueOf(referenceStruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		expectedKeys[val.Type().Field(i).Name] = true
	}

	for key, value := range rawReq {
		// Check if key is expected
		if _, exists := expectedKeys[key]; !exists {
			return customerrors.ErrValidateKeysAgainstStruct(nil, fmt.Sprintf("Unexpected key '%s' in the request", key))
		}

		// Recursively check nested structures
		if nestedMap, ok := value.(map[string]interface{}); ok {
			field, _ := val.Type().FieldByName(key)
			if field.Type.Kind() == reflect.Struct {
				if err := ValidateKeysAgainstStruct(nestedMap, reflect.New(field.Type).Interface()); err != nil {
					return customerrors.ErrValidateKeysAgainstStruct(nil, fmt.Sprintf("key '%s': %v", key, err))
				}
			}
		} else if nestedSlice, ok := value.([]interface{}); ok {
			field, _ := val.Type().FieldByName(key)
			if field.Type.Elem().Kind() == reflect.Struct {
				for i, nestedElement := range nestedSlice {
					if nestedMap, ok := nestedElement.(map[string]interface{}); ok {
						if err := ValidateKeysAgainstStruct(nestedMap, reflect.New(field.Type.Elem()).Interface()); err != nil {
							return customerrors.ErrValidateKeysAgainstStruct(nil, fmt.Sprintf("key '%s' index %d: %v", key, i, err))
						}
					}
				}
			}
		}
	}
	return nil
}
