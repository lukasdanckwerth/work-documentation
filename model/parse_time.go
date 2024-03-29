package model

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseTime parses a time string
//
// Supported formats:
// 15m
// 1h
// 1:30h
// 1:30
// 1.5h
// 1.5
func ParseTime(input string) (int64, error) {

	if strings.Contains(input, ":") {

		if strings.HasSuffix(input, "h") {
			input = input[:len(input)-1]
		}

		var split = strings.Split(input, ":")
		var hours, error = strconv.ParseInt(split[0], 10, 64)

		if error != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		var minutes, error1 = strconv.ParseInt(split[1], 10, 64)

		if error1 != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		return (hours * 60) + minutes, nil
	} else if strings.Contains(input, ".") {

		if strings.HasSuffix(input, "h") {
			input = input[:len(input)-1]
		}

		var split = strings.Split(input, ".")
		var hours, error = strconv.ParseInt(split[0], 10, 64)

		if error != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		var fractial, error1 = strconv.ParseFloat("0."+split[1], 64)

		if error1 != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		var minutes = int64(fractial * 60)

		return (hours * 60) + minutes, nil
	} else if strings.HasSuffix(input, "h") {

		var hoursString = input[:len(input)-1]
		var hours, error = strconv.ParseInt(hoursString, 10, 64)

		if error != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		return hours * 60, nil
	} else if strings.HasSuffix(input, "m") {

		var minutesString = input[:len(input)-1]
		var minutes, error = strconv.ParseInt(minutesString, 10, 64)

		if error != nil {
			return 0, fmt.Errorf("invalid time format: %v", input)
		}

		return minutes, nil
	} else {

		return 0, fmt.Errorf("invalid time format: %v", input)
	}
}
