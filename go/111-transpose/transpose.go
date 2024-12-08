package transpose

func Transpose(input []string) []string {
	result := []string{}
	for r, row := range input {
		for c, char := range row {
			for len(result) <= c {
				result = append(result, "")
			}

			for len(result[c]) < r {
				result[c] += " "
			}

			result[c] += string(char)
		}
	}
	return result
}

//func Transpose(input []string) []string {
//	result := []string{}
//
//	for r, row := range input {
//		temp := make([][]string, len(row))
//		for j, cell := range row {
//			str := string(cell)
//			temp[j] = append(temp[j], str)
//		}
//
//		for i, slc := range temp {
//			if len(result) <= i {
//				finalString := strings.Join(slc, "")
//				if r != 0 {
//					for i := 0; i < r; i++ {
//						finalString = " " + finalString
//					}
//				}
//				result = append(result, finalString)
//			} else {
//				if r != len(result[i]) {
//					result[i] += " "
//				}
//				result[i] += strings.Join(slc, "")
//			}
//		}
//	}
//
//	return result
//}
