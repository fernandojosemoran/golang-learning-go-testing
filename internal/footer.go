package internal

import "strconv"

func Fooer(input int) string {
	var isFood bool = (input % 3) == 0
	if isFood {
		return "Food"
	}

	return strconv.Itoa(input)
}
