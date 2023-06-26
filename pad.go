package pad

import "strings"

// Pads a given string to a given length
// or truncates the string if it is greater than that length
func Pad(s string, max uint) string {
	length := uint(len(s))
	if length > max {
		return s[:max-1]
	}
	s += strings.Repeat(" ", int(max-length))
	return s
}