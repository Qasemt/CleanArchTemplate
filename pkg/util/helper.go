package util

import "strconv"

func ToUint(str string) (uint, error) {
	uintValue, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(uintValue), nil
}
func ToInt(str string) (int, error) {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}
