package day02

import (
	"strings"

	. "github.com/lucascott/advent-of-code-2025/lib"
)

type Range struct {
	start int
	end   int
}

func parseRanges(input string) []Range {
	rangesStr := strings.Split(input, ",")
	var ranges []Range
	for _, rangeStr := range rangesStr {
		range_ := strings.SplitN(rangeStr, "-", 2)
		ranges = append(ranges, Range{
			ParseInt(range_[0]),
			ParseInt(range_[1]),
		})
	}
	return ranges
}
