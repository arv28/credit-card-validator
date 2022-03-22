package lib

import (
	"fmt"
	"strconv"
	"unicode"
)

type digits [4]int

// at returns the digits from the start to the given length
func (d *digits) at(i int) int {
	return d[i-1]
}

// check if credit card number is valid
func (c *Card) Validate() bool {
	isSecond := false
	nSum := 0

	cardLength := len(c.Number)

	if cardLength == 0 || cardLength < 12 || cardLength > 19 {
		return false
	}

	// length of card number
	nDigits := len(c.Number)

	for i := nDigits - 1; i >= 0; i-- {
		char := string(c.Number[i])
		if char == " " {
			continue
		}
		num, err := strconv.Atoi(char)
		if err != nil {
			fmt.Printf("Invalid card number %s\n", c.Number)
			return false
		}

		if isSecond {
			num = num * 2
		}

		// add two digits to handle cases that make two digits after doubling
		nSum += num / 10
		nSum += num % 10

		isSecond = !isSecond
	}

	if nSum%10 == 0 {
		return true
	}

	return false
}

// GetScheme returns the card type the credit card has
func (c *Card) GetScheme() (cardType, error) {
	digitLen := getDigitLen(c.Number)
	ccDigits := digits{}

	for i := 0; i < 4; i++ {
		if i < digitLen {
			ccDigits[i], _ = strconv.Atoi(c.Number[:i+1])
		}
	}

	// switch compares the digits of the card
	switch {
	// American Express	scheme check
	case ccDigits.at(2) == 34 || ccDigits.at(2) == 37 && digitLen == 15:
		return AmericanExpress, nil

	// JCB scheme check
	case ccDigits.at(4) >= 3528 && ccDigits.at(4) <= 3589 && (digitLen > 15 && digitLen < 20):
		return JCB, nil

	// Maestro scheme check
	case ccDigits.at(1) == 6 || ccDigits.at(2) == 50 || (ccDigits.at(2) >= 56 && ccDigits.at(2) <= 58) &&
		(digitLen > 11 && digitLen < 20):
		return Maestro, nil

	// Visa scheme check
	case ccDigits.at(1) == 4 && (digitLen == 13 || digitLen == 16 || digitLen == 19):
		return Visa, nil

	// MasterCard scheme check
	case (ccDigits.at(2) >= 51 && ccDigits.at(2) <= 55) || (ccDigits.at(4) >= 2221 && ccDigits.at(4) <= 2720) &&
		(digitLen == 16):
		return Mastercard, nil

	default:
		return -1, fmt.Errorf("unknown creditcard type")
	}

}

// get the digits length in a card number
func getDigitLen(number string) int {
	count := 0

	for i := 0; i < len(number); i++ {
		if unicode.IsDigit(rune(number[i])) {
			count++
		}
	}

	return count
}
