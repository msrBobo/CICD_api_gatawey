package v1

import (
	"fmt"
	"strconv"
)

func ParseQueryParams(page string, limit string) (uint64, uint64, error) {
	if len(page) < 1 && len(limit) < 1 {
		return 1, 10, nil
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return 0, 0, err
	}
	if pageInt < 0 {
		return 0, 0, fmt.Errorf("page cannot be a negative number")
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return 0, 0, err
	}
	if limitInt < 0 {
		return 0, 0, fmt.Errorf("limit cannot be a negative number")
	}
	return uint64(pageInt), uint64(limitInt), nil
}
