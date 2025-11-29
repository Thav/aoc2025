package lists

import (
	"fmt"
	"strconv"
	"strings"
)

func ImportRowLists(b []byte, split string) (lists [][]string) {
	input := string(b)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := strings.Split(line, split)
		lists = append(lists, row)
	}
	return
}
func ImportRowListsInt(b []byte, split string) (lists [][]int) {
	strLists := ImportRowLists(b, split)
	for _, list := range strLists {
		lists = append(lists, StringSliceToIntSlice(list))
	}
	return
}

func ImportLeftRightLists(b []byte, split string) (left, right []string) {
	input := string(b)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		values := strings.Split(line, split)
		if len(values) != 2 {
			panic(fmt.Sprintln("There were not two values in this line, ", line))
		}
		left = append(left, values[0])
		right = append(right, values[1])
	}
	return
}

func ImportLeftRightListsInt(b []byte, split string) (left, right []int) {
	l, r := ImportLeftRightLists(b, split)
	left = StringSliceToIntSlice(l)
	right = StringSliceToIntSlice(r)
	return
}

func StringSliceToIntSlice(strings []string) (ints []int) {
	for _, strValue := range strings {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			panic(fmt.Sprintln("Couldn't convert ", strValue, " to an integer"))
		}
		ints = append(ints, value)
	}
	return
}
