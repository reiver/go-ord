package orden

import (
	"strconv"
)

// FormatInt64 return the number in its numeric- ordinal format.
//
// For example —
//
//	0 → "0th"
//
//	1 → "1st"
//
//	2 → "2nd"
//
//	3 → "3rd"
//
//	4 → "4th"
//
//	5 → "5th"
//
//	...
//
//	10 → "10th"
//
//	11 → "11th"
//
//	12 → "12th"
//
//	13 → "13th"
//
//	14 → "14th"
//
//	...
//
//	1000 → "1000th"
//
//	1001 → "1001st"
//
//	1002 → "1002nd"
//
//	1003 → "1003rd"
//
//	1004 → "1004th"
//
//	...
func FormatInt64(n int64) string {

	var p []byte
	{
		p = append(p, strconv.FormatInt(n, 10)...)

		p = append(p, suffixInt64(n)...)
	}

	return string(p)
}
