package convert

import "strconv"

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
