package lib

type Card struct {
	// Number is the credit card number
	Number string
}

type cardType int

const (
	// Unknown card type
	Unknown cardType = iota
	// AmericanExpress  card type
	AmericanExpress
	// JCB card type
	JCB
	// Maestro card type
	Maestro
	// Mastercard card type
	Mastercard
	// Visa card type
	Visa
)
