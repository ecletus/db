package utils

import (
	"strings"
	"reflect"
	"time"
)

func TrimStrings(fields ... *string) {
	for _, f := range fields {
		*f = strings.TrimSpace(*f)
	}
}

func SetTimeValue(obj reflect.Value, fieldName string, value *time.Time, check func(v *time.Time) bool) bool {
	if f := obj.FieldByName(fieldName); f.IsValid() {
		fv := f.Interface()
		if v, ok := fv.(*time.Time); ok {
			if check == nil || check(v) {
				f.Set(reflect.ValueOf(value))
				return true
			}
		} else if v, ok := fv.(time.Time); ok {
			if check == nil || check(&v) {
				f.Set(reflect.ValueOf(*value))
				return true
			}
		}
	}
	return false
}

func SetTimeValueIfEmpty(obj reflect.Value, fieldName string, value *time.Time) bool {
	return SetTimeValue(obj, fieldName, value, func(v *time.Time) bool {
		return v == nil || v.IsZero()
	})
}

func SetStringValue(obj reflect.Value, fieldName string, value *string, check func(v *string) bool) bool {
	if f := obj.FieldByName(fieldName); f.IsValid() {
		fv := f.Interface()
		if v, ok := fv.(*string); ok {
			if check == nil || check(v) {
				f.Set(reflect.ValueOf(value))
				return true
			}
		} else if v, ok := fv.(string); ok {
			if check == nil || check(&v) {
				if value == nil {
					f.SetString("")
				} else {
					f.SetString(*value)
				}
				return true
			}
		}
	}
	return false
}

func SetStringValueIfEmpty(obj reflect.Value, fieldName string, value *string) bool {
	return SetStringValue(obj, fieldName, value, func(v *string) bool {
		return v == nil || *v == ""
	})
}