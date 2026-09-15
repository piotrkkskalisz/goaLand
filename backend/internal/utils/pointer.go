package utils

func PointerToInt(value int) *int {
	return &value
}

func IntOrZero(value *int) int {
	if value == nil {
		return 0
	}
	return *value
}
