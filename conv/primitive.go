package conv

import (
	"fmt"
	"strconv"
)

func MustAtoi64(s string) int64 {
	if n, err := strconv.Atoi(s); err == nil {
		return int64(n)
	} else {
		panic(fmt.Sprintf("Failed to convert string %s to number", s))
	}
}

func MustAtoi(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	} else {
		panic(fmt.Sprintf("Failed to convert string %s to number", s))
	}
}

func MustParseOctValue(s string) int {
	if val, err := strconv.ParseInt(s, 0, 32); err == nil {
		return int(val)
	} else {
		panic(fmt.Sprintf("Failed to parse string %s to float", s))
	}
}

func MustParseFloat32(s string) float32 {
	if val, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(val)
	} else {
		panic(fmt.Sprintf("Failed to parse string %s to float", s))
	}
}
