package gopherriz

// Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

// If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
// If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
// Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
// Given an integer, convert it to a Roman numeral.

// our solution

func intToRoman(num int) string {
	roman := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	values := []int{
		1000, 900, 500, 400, 100, 90,
		50, 40, 10, 9, 5, 4, 1,
	}

	numCpy := num
	i := 0
	romanNum := ""

	for numCpy > 0 {
		if numCpy < values[i] {
			i = i + 1
			continue
		}
		numCpy = numCpy - values[i]
		romanNum = romanNum + roman[values[i]]
		i = 0
	}
	return romanNum
}

// extremely cheap and simple alternative and works exceptionally well withing a range of 1-3999

var (
	r0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	r1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	r2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	r3 = []string{"", "M", "MM", "MMM"}
)

func intToRomanOne(num int) string {
	// This is efficient in Go. The 4 operands are evaluated,
	// then a single allocation is made of the exact size needed for the result.
	return r3[num%1e4/1e3] + r2[num%1e3/1e2] + r1[num%100/10] + r0[num%10]
}
