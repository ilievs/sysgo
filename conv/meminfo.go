package conv

import (
	"strconv"
	"strings"
)

func MustConvertKbValueToInt(kbValue string) int {
	parts := strings.Split(kbValue, " ")
	if len(parts) != 2 {
		panic("Cannot extract KB, got: " + kbValue)
	}

	if val, err := strconv.Atoi(parts[0]); err == nil {
		return val
	} else {
		panic("Cannot extract KB, got: " + kbValue)
	}
}
