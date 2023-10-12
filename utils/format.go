package utils

import "strconv"

func ConvertIntToString(value int) string {
	return strconv.Itoa(value)
}

func ConvertStringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}
