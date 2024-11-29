// Package cmd
/*
Copyright © 2024 leig <leigme@gmail.com>
*/
package cmd

import "cmp"

func Unique[T cmp.Ordered](ss []T) []T {
	size := len(ss)
	if size == 0 {
		return []T{}
	}
	newSlices := make([]T, 0)
	mapSlices := make(map[T]byte)
	for _, v := range ss {
		if _, ok := mapSlices[v]; !ok {
			mapSlices[v] = 1
			newSlices = append(newSlices, v)
		}
	}
	return newSlices
}
