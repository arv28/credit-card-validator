# credit-card-validator
Verify the correctness of a supplied credit card number and determine its card scheme

## Supported Functions
- `Validate()`- Validates the credit card number using Luhn's algorithm
- `GetScheme()` - Determines the card type like Visa, MasterCard etc.

## Example Usage
```
c := &Card{
    Number: "1234567890",
}

// validate
c.Validate()

// get scheme
type, err := c.GetScheme()
```

## Tests
To run the tests, simply run the following commands in terminal:
```
go test -v ./...
```