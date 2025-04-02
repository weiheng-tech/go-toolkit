package utils

func Contain[T Number](n T, s []T) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}
