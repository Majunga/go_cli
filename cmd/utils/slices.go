package utils

import "slices"

func Reduce(slice []string, predicate func(s string) bool) []string {
	i := slices.IndexFunc(slice, predicate)

	if i < 0 {
		return slice
	}

	return Reduce(slices.DeleteFunc(slice, predicate), predicate)
}
