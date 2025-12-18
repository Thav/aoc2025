package convert

import (
	"log"
	"strconv"
)

func StringSliceToIntSlice(input []string) (output []int, err error) {
	for i := range len(input) {
		integer, err := strconv.Atoi(input[i])
		if err != nil {
			var empty []int
			return empty, err
		}
		output = append(output, integer)
	}
	return output, nil
}

func ToInt(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln("couldn't convert to int: ", s)
	}
	return value
}

func SliceToInt(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i] = ToInt(s)
	}
	return ints
}
