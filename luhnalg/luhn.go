package luhnalg

import "strconv"

type LuhnAlgorithm struct {

}

func NewLuhnAlgorithm() *LuhnAlgorithm{
	return &LuhnAlgorithm{}
}

func (this *LuhnAlgorithm) ComputeCheckDigit(number string) int{
	if len(number) == 0{
		return 0
	}

	sum := 0
	numberOfDigits := len(number)
	shouldMultiplyDigitByTwo := true

	for i:= numberOfDigits-1; i>=0; i-- {
		currentDigit, _ := strconv.Atoi(string(number[i]))
		if shouldMultiplyDigitByTwo {
			currentDigit *= 2
		}

		if isTwoDigit(currentDigit) {
			currentDigit = getFirstDigit(currentDigit) + getSecondDigit(currentDigit)
		}

		sum += currentDigit
		shouldMultiplyDigitByTwo = !shouldMultiplyDigitByTwo
	}

	checkDigit := sum%10
	return checkDigit

}

func isTwoDigit(n int) bool{
	return n>9 && n<100
}

func getFirstDigit(n int) int{
	return n/10
}

func getSecondDigit(n int) int{
	return n%10
}
