package tool

func Index(str string, data []string) int {
	for index, value := range data {
		if str == value {
			return index
		}
	}
	return -1 //not found.
}
