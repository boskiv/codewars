package roman

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)


func Decode(roman string) int {
	var stringToRoman = map[string]int{
		"I": I,
		"V": V,
		"X": X,
		"L": L,
		"C": C,
		"D": D,
		"M": M,
	}

	var result, number, previousNumber int
	for i, char := range roman {
		letter := string(char)
		number = stringToRoman[letter]

		if len(roman) == 1 {
			return stringToRoman[letter]
		}
		if i == 0 {
			result = number
			previousNumber = number
			continue
		}

		if previousNumber >= number {
			result = result + number
		} else {
			result = result + number - previousNumber * 2
		}

		previousNumber = number

	}

	return result
}
