package flatten

func Flatten(input interface{}) []interface{} {
	flat := []interface{}{}

	switch i := input.(type) {
	case []interface{}:
		for _, val := range i {
			flat = append(flat, Flatten(val)...)
		}
	case interface{}:
		flat = append(flat, i)
	}

	return flat
}

//func Flatten(nested interface{}) []interface{} {
//	var output []interface{} = []interface{}{}
//	val := reflect.ValueOf(nested)
//
//	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
//		for i := 0; i < val.Len(); i++ {
//			item := val.Index(i).Interface()
//			output = append(output, Flatten(item)...)
//		}
//	} else {
//		if nested != nil {
//			output = append(output, nested)
//
//		}
//	}
//
//	return output
//}
