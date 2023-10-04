package utils

import (
	"reflect"

	"github.com/FUJI0130/curriculum/src/core/utils/customerrors"
)

func StructToMap(req interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	val := reflect.ValueOf(req)
	if val.Kind() != reflect.Ptr {
		return nil, customerrors.ErrStructToMap(nil, "Expected a pointer but got "+val.Kind().String())
	}
	val = val.Elem() // Now it's safe to call Elem
	typ := val.Type()

	// Iterate over struct fields
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Ignore unexported fields
		if fieldType.PkgPath != "" {
			continue
		}

		key := fieldType.Name

		// Recursively process nested structures
		if field.Kind() == reflect.Struct {
			nestedMap, err := StructToMap(field.Addr().Interface())
			if err != nil {
				return nil, err
			}
			result[key] = nestedMap
		} else if field.Kind() == reflect.Slice {
			var sliceData []interface{}
			for si := 0; si < field.Len(); si++ {
				sliceElem := field.Index(si)
				if sliceElem.Kind() == reflect.Struct {
					nestedMap, err := StructToMap(sliceElem.Addr().Interface())
					if err != nil {
						return nil, err
					}
					sliceData = append(sliceData, nestedMap)
				} else {
					sliceData = append(sliceData, sliceElem.Interface())
				}
			}
			result[key] = sliceData
		} else {
			// Plain field, just set the value
			result[key] = field.Interface()
		}
	}
	return result, nil
}
