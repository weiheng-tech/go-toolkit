package slice

func Contain[T comparable](n T, s []T) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}
	return false
}

func Concat[T any](slices ...[]T) []T {
	totalLen := 0
	for _, v := range slices {
		totalLen += len(v)
		if totalLen < 0 {
			panic("len out of range")
		}
	}
	result := make([]T, 0, totalLen)

	for _, v := range slices {
		result = append(result, v...)
	}

	return result
}

func Difference[T comparable](slice, comparedSlice []T) []T {
	var result []T

	if len(slice) == 0 {
		return result
	}

	comparedMap := make(map[T]struct{}, len(comparedSlice))
	for _, v := range comparedSlice {
		comparedMap[v] = struct{}{}
	}

	for _, v := range slice {
		if _, found := comparedMap[v]; !found {
			result = append(result, v)
		}
	}

	return result
}

func Equal[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
}
