package orden

import (
	"testing"
)

func TestSuffixInt64(t *testing.T) {

	tests := []struct{
		Number int64
		Expected string
	}{
		{
			Number:   0,
			Expected: "th",
		},
		{
			Number:   1,
			Expected: "st",
		},
		{
			Number:   2,
			Expected: "nd",
		},
		{
			Number:   3,
			Expected: "rd",
		},
		{
			Number:   4,
			Expected: "th",
		},
		{
			Number:   5,
			Expected: "th",
		},
		{
			Number:   6,
			Expected: "th",
		},
		{
			Number:   7,
			Expected: "th",
		},
		{
			Number:   8,
			Expected: "th",
		},
		{
			Number:   9,
			Expected: "th",
		},



		{
			Number:  10,
			Expected: "th",
		},
		{
			Number:  11,
			Expected: "th",
		},
		{
			Number:  12,
			Expected: "th",
		},
		{
			Number:  13,
			Expected: "th",
		},
		{
			Number:  14,
			Expected: "th",
		},
		{
			Number:  15,
			Expected: "th",
		},
		{
			Number:  16,
			Expected: "th",
		},
		{
			Number:  17,
			Expected: "th",
		},
		{
			Number:  18,
			Expected: "th",
		},
		{
			Number:  19,
			Expected: "th",
		},



		{
			Number:  20,
			Expected: "th",
		},
		{
			Number:  21,
			Expected: "st",
		},
		{
			Number:  22,
			Expected: "nd",
		},
		{
			Number:  23,
			Expected: "rd",
		},
		{
			Number:  24,
			Expected: "th",
		},
		{
			Number:  25,
			Expected: "th",
		},
		{
			Number:  26,
			Expected: "th",
		},
		{
			Number:  27,
			Expected: "th",
		},
		{
			Number:  28,
			Expected: "th",
		},
		{
			Number:  29,
			Expected: "th",
		},



		{
			Number:  30,
			Expected: "th",
		},
		{
			Number:  31,
			Expected: "st",
		},
		{
			Number:  32,
			Expected: "nd",
		},
		{
			Number:  33,
			Expected: "rd",
		},
		{
			Number:  34,
			Expected: "th",
		},
		{
			Number:  35,
			Expected: "th",
		},
		{
			Number:  36,
			Expected: "th",
		},
		{
			Number:  37,
			Expected: "th",
		},
		{
			Number:  38,
			Expected: "th",
		},
		{
			Number:  39,
			Expected: "th",
		},



		{
			Number:  40,
			Expected: "th",
		},
		{
			Number:  41,
			Expected: "st",
		},
		{
			Number:  42,
			Expected: "nd",
		},
		{
			Number:  43,
			Expected: "rd",
		},
		{
			Number:  44,
			Expected: "th",
		},
		{
			Number:  45,
			Expected: "th",
		},
		{
			Number:  46,
			Expected: "th",
		},
		{
			Number:  47,
			Expected: "th",
		},
		{
			Number:  48,
			Expected: "th",
		},
		{
			Number:  49,
			Expected: "th",
		},



		{
			Number:  50,
			Expected: "th",
		},
		{
			Number:  51,
			Expected: "st",
		},
		{
			Number:  52,
			Expected: "nd",
		},
		{
			Number:  53,
			Expected: "rd",
		},
		{
			Number:  54,
			Expected: "th",
		},
		{
			Number:  55,
			Expected: "th",
		},
		{
			Number:  56,
			Expected: "th",
		},
		{
			Number:  57,
			Expected: "th",
		},
		{
			Number:  58,
			Expected: "th",
		},
		{
			Number:  59,
			Expected: "th",
		},



		{
			Number:  60,
			Expected: "th",
		},
		{
			Number:  61,
			Expected: "st",
		},
		{
			Number:  62,
			Expected: "nd",
		},
		{
			Number:  63,
			Expected: "rd",
		},
		{
			Number:  64,
			Expected: "th",
		},
		{
			Number:  65,
			Expected: "th",
		},
		{
			Number:  66,
			Expected: "th",
		},
		{
			Number:  67,
			Expected: "th",
		},
		{
			Number:  68,
			Expected: "th",
		},
		{
			Number:  69,
			Expected: "th",
		},



		{
			Number:  70,
			Expected: "th",
		},
		{
			Number:  71,
			Expected: "st",
		},
		{
			Number:  72,
			Expected: "nd",
		},
		{
			Number:  73,
			Expected: "rd",
		},
		{
			Number:  74,
			Expected: "th",
		},
		{
			Number:  75,
			Expected: "th",
		},
		{
			Number:  76,
			Expected: "th",
		},
		{
			Number:  77,
			Expected: "th",
		},
		{
			Number:  78,
			Expected: "th",
		},
		{
			Number:  79,
			Expected: "th",
		},



		{
			Number:  80,
			Expected: "th",
		},
		{
			Number:  81,
			Expected: "st",
		},
		{
			Number:  82,
			Expected: "nd",
		},
		{
			Number:  83,
			Expected: "rd",
		},
		{
			Number:  84,
			Expected: "th",
		},
		{
			Number:  85,
			Expected: "th",
		},
		{
			Number:  86,
			Expected: "th",
		},
		{
			Number:  87,
			Expected: "th",
		},
		{
			Number:  88,
			Expected: "th",
		},
		{
			Number:  89,
			Expected: "th",
		},



		{
			Number:  90,
			Expected: "th",
		},
		{
			Number:  91,
			Expected: "st",
		},
		{
			Number:  92,
			Expected: "nd",
		},
		{
			Number:  93,
			Expected: "rd",
		},
		{
			Number:  94,
			Expected: "th",
		},
		{
			Number:  95,
			Expected: "th",
		},
		{
			Number:  96,
			Expected: "th",
		},
		{
			Number:  97,
			Expected: "th",
		},
		{
			Number:  98,
			Expected: "th",
		},
		{
			Number:  99,
			Expected: "th",
		},



		{
			Number: 100,
			Expected: "th",
		},
		{
			Number: 101,
			Expected: "st",
		},
		{
			Number: 102,
			Expected: "nd",
		},
		{
			Number: 103,
			Expected: "rd",
		},
		{
			Number: 104,
			Expected: "th",
		},
		{
			Number: 105,
			Expected: "th",
		},
		{
			Number: 106,
			Expected: "th",
		},
		{
			Number: 107,
			Expected: "th",
		},
		{
			Number: 108,
			Expected: "th",
		},
		{
			Number: 109,
			Expected: "th",
		},



		{
			Number: 110,
			Expected: "th",
		},
		{
			Number: 111,
			Expected: "th",
		},
		{
			Number: 112,
			Expected: "th",
		},
		{
			Number: 113,
			Expected: "th",
		},
		{
			Number: 114,
			Expected: "th",
		},
		{
			Number: 115,
			Expected: "th",
		},
		{
			Number: 116,
			Expected: "th",
		},
		{
			Number: 117,
			Expected: "th",
		},
		{
			Number: 118,
			Expected: "th",
		},
		{
			Number: 119,
			Expected: "th",
		},



		{
			Number: 120,
			Expected: "th",
		},
		{
			Number: 121,
			Expected: "st",
		},
		{
			Number: 122,
			Expected: "nd",
		},
		{
			Number: 123,
			Expected: "rd",
		},
		{
			Number: 124,
			Expected: "th",
		},
		{
			Number: 125,
			Expected: "th",
		},
		{
			Number: 126,
			Expected: "th",
		},
		{
			Number: 127,
			Expected: "th",
		},
		{
			Number: 128,
			Expected: "th",
		},
		{
			Number: 129,
			Expected: "th",
		},



		{
			Number: 130,
			Expected: "th",
		},
		{
			Number: 131,
			Expected: "st",
		},
		{
			Number: 132,
			Expected: "nd",
		},
		{
			Number: 133,
			Expected: "rd",
		},
		{
			Number: 134,
			Expected: "th",
		},
		{
			Number: 135,
			Expected: "th",
		},
		{
			Number: 136,
			Expected: "th",
		},
		{
			Number: 137,
			Expected: "th",
		},
		{
			Number: 138,
			Expected: "th",
		},
		{
			Number: 139,
			Expected: "th",
		},



		{
			Number: 140,
			Expected: "th",
		},
		{
			Number: 141,
			Expected: "st",
		},
		{
			Number: 142,
			Expected: "nd",
		},
		{
			Number: 143,
			Expected: "rd",
		},
		{
			Number: 144,
			Expected: "th",
		},
		{
			Number: 145,
			Expected: "th",
		},
		{
			Number: 146,
			Expected: "th",
		},
		{
			Number: 147,
			Expected: "th",
		},
		{
			Number: 148,
			Expected: "th",
		},
		{
			Number: 149,
			Expected: "th",
		},



		{
			Number: 150,
			Expected: "th",
		},
		{
			Number: 151,
			Expected: "st",
		},
		{
			Number: 152,
			Expected: "nd",
		},
		{
			Number: 153,
			Expected: "rd",
		},
		{
			Number: 154,
			Expected: "th",
		},
		{
			Number: 155,
			Expected: "th",
		},
		{
			Number: 156,
			Expected: "th",
		},
		{
			Number: 157,
			Expected: "th",
		},
		{
			Number: 158,
			Expected: "th",
		},
		{
			Number: 159,
			Expected: "th",
		},






		{
			Number: 1000,
			Expected:  "th",
		},
		{
			Number: 1001,
			Expected:  "st",
		},
		{
			Number: 1002,
			Expected:  "nd",
		},
		{
			Number: 1003,
			Expected:  "rd",
		},
		{
			Number: 1004,
			Expected:  "th",
		},
		{
			Number: 1005,
			Expected:  "th",
		},
		{
			Number: 1006,
			Expected:  "th",
		},
		{
			Number: 1007,
			Expected:  "th",
		},
		{
			Number: 1008,
			Expected:  "th",
		},
		{
			Number: 1009,
			Expected:  "th",
		},



		{
			Number: 1010,
			Expected:  "th",
		},
		{
			Number: 1011,
			Expected:  "th",
		},
		{
			Number: 1012,
			Expected:  "th",
		},
		{
			Number: 1013,
			Expected:  "th",
		},
		{
			Number: 1014,
			Expected:  "th",
		},
		{
			Number: 1015,
			Expected:  "th",
		},
		{
			Number: 1016,
			Expected:  "th",
		},
		{
			Number: 1017,
			Expected:  "th",
		},
		{
			Number: 1018,
			Expected:  "th",
		},
		{
			Number: 1019,
			Expected:  "th",
		},



		{
			Number: 1020,
			Expected:  "th",
		},
		{
			Number: 1021,
			Expected:  "st",
		},
		{
			Number: 1022,
			Expected:  "nd",
		},
		{
			Number: 1023,
			Expected:  "rd",
		},
		{
			Number: 1024,
			Expected:  "th",
		},
		{
			Number: 1025,
			Expected:  "th",
		},
		{
			Number: 1026,
			Expected:  "th",
		},
		{
			Number: 1027,
			Expected:  "th",
		},
		{
			Number: 1028,
			Expected:  "th",
		},
		{
			Number: 1029,
			Expected:  "th",
		},
	}

	for testNumber, test := range tests {

		var expected string = test.Expected
		var actual   string = suffixInt64(test.Number)

		if expected != actual {
			t.Errorf("For test #%d, the actual suffix value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("NUMBER: %d", test.Number)
			continue
		}
	}
}
