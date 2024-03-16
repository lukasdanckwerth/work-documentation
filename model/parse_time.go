package model

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseTime(input string) (int64, error) {

	if strings.HasSuffix(input, "h") {

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
