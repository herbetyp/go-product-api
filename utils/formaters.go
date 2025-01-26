package utils

import "strconv"

func StringToUint(id string) (uint, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(idInt), nil
}
