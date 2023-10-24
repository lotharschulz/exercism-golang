package armstrong

import (
    "math"
)

// determines if a given number is an Armstrong number
// see also https://github.com/lotharschulz/armstrongNumbers/
func IsNumber(n int) bool {
	if n < 0 {
        return false
    }
	digits := getDigits(n)
    numberOfDigits := getNumberOfDigits(digits)
    total := 0
	for _, v := range digits{
        total = total + int(math.Pow(float64(v),float64(numberOfDigits)))
    }
    if total == n{
        return true
    }
    return false

}

func getDigits(n int) []int{
    digits := []int{}
    for n !=0 {
        digits = append(digits, n % 10)
        n /= 10
    }
    reverseInt(digits)
    return digits
}

func reverseInt(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
            s[i], s[j] = s[j], s[i]
    }
}

func getNumberOfDigits(digits []int) int{
	return len(digits)
}
