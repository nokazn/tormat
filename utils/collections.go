package utils

func Unshift[T any](slice []T, v T) []T {
	slice = append(slice[:1], slice[0:]...)
	slice[0] = v
	return slice
}

type callback[I any, O any] func(input I, index uint) O

func Map[I any, O any](list []I, callback callback[I, O]) []O {
	l := make([]O, len(list))
	for i, v := range list {
		l[i] = callback(v, uint(i))
	}
	return l
}
