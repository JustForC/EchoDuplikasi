package helpers

import "strconv"

func ConvertId(id string) int {
	converted, _ := strconv.Atoi(id)

	return converted
}
