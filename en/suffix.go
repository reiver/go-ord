package orden

func suffixInt64(n int64) string {

	{
		d := n % 100

		switch d {
		case 11, 12, 13:
			return "th"
		}
	}

	{
		d := n % 10

		switch d {
		case 1:
			return "st"
		case 2:
			return "nd"
		case 3:
			return "rd"
		default:
			return "th"
		}
	}
}
