package utils

import "errors"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Max[T Number](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, errors.New("cannot find max of empty slice")
	}

	maxNumber := s[0]
	for _, v := range s[1:] {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber, nil
}

func Min[T Number](s []T) (T, error) {
	if len(s) == 0 {
		var zero T
		return zero, errors.New("cannot find min of empty slice")
	}

	minNumber := s[0]
	for _, v := range s[1:] {
		if v < minNumber {
			minNumber = v
		}
	}
	return minNumber, nil
}
