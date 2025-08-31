package utils

import (
    "reflect"
)

// StructToMap converts a struct with pointer fields into a map for updates
func StructToMap(s interface{}) map[string]interface{} {
    updates := make(map[string]interface{})
    val := reflect.ValueOf(s).Elem() // assuming pointer to struct
    typ := val.Type()

    for i := 0; i < val.NumField(); i++ {
        field := val.Field(i)
        fieldType := typ.Field(i)

        // Skip nil pointers
        if field.Kind() == reflect.Ptr && !field.IsNil() {
            jsonTag := fieldType.Tag.Get("json")
            if jsonTag == "" {
                jsonTag = fieldType.Name
            }
            updates[jsonTag] = field.Elem().Interface()
        }
    }

    return updates
}
