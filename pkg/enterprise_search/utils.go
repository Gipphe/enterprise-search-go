package enterprise_search

import (
	"net/url"
	"reflect"
	"strconv"
)

func paramsToValues(params interface{}) url.Values {
	values := url.Values{}
	if params == nil {
		return values
	}

	v := reflect.ValueOf(params)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return values
	}

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
			case reflect.Slice:
				if ss, ok := field.Interface().([]string); ok {
					for _, s := range ss {
						values.Add(tag, s)
					}
				}
			case reflect.Ptr:
				if b, ok := field.Interface().(*bool); ok {
					values.Set(tag, strconv.FormatBool(*b))
				}
			}
		}
	}

	return values
}