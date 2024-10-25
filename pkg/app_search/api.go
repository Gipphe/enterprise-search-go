package app_search

import (
	"net/url"
	"reflect"
	"strconv"
)

// paramsToValues converts a struct to url.Values
func paramsToValues(params interface{}) url.Values {
	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}

	values := url.Values{}
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" || tag == "-" {
			continue
		}

		if !field.IsZero() {
			switch field.Kind() {
			case reflect.Ptr:
				if field.Elem().Kind() == reflect.Int {
					values.Set(tag, strconv.Itoa(int(field.Elem().Int())))
				}
			case reflect.Slice:
				for _, id := range field.Interface().([]string) {
					values.Add(tag, id)
				}
			}
		}
	}

	return values
}
