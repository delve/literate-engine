package stringy

import (
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// PrefixOrDie dies if the input value does not begin with prefix
func PrefixOrDie(prefix, value string) {
	if !strings.HasPrefix(value, prefix) {
		logrus.Fatalf("prefix %s missing from: %s", prefix, value)
	}
}

// Atoi trims space from its input string and converts it to an int. or dies trying.
func Atoi(s string) int {
	ret, err := strconv.Atoi(strings.TrimSpace(s))
	Check(err)
	return ret
}

// MustFloat trims space from its input string and converts it to float64. or dies trying.
func MustFloat(s string) float64 {
	ret, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	Check(err)
	return ret
}