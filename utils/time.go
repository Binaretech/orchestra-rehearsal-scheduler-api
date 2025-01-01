package utils

import (
	"regexp"
	"strconv"
)

var tzOffsetRegex = regexp.MustCompile(`^([+-]?)(0?\d|1[0-4]):([0-5]\d)$`)

func ValidateTimeZoneOffset(offset string) bool {
	matches := tzOffsetRegex.FindStringSubmatch(offset)
	if len(matches) == 0 {
		return false
	}

	signStr := matches[1]
	hourStr := matches[2]
	minuteStr := matches[3]

	hours, errH := strconv.Atoi(hourStr)

	if errH != nil {
		return false
	}

	if _, errM := strconv.Atoi(minuteStr); errM != nil {
		return false
	}

	if signStr == "-" {
		hours = -hours
	}

	if hours < -14 || hours > 14 {
		return false
	}

	return true
}
