package lib

import (
	"testing"
)

func TestValidate(t *testing.T) {
	type args struct {
		cardNumber string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{
			name: "valid card",
			args: args{
				cardNumber: "3714 4963 539 8431",
			},
			want: true,
		},
		{
			name: "Invalid card",
			args: args{
				cardNumber: "3714496353984312",
			},
			want: false,
		},
		{
			name: "empty input",
			args: args{
				cardNumber: "",
			},
			want: false,
		},
		{
			name: "Invalid input",
			args: args{
				cardNumber: "1234abc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Number: tt.args.cardNumber,
			}
			if got := c.Validate(); got != tt.want {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_GetScheme(t *testing.T) {
	tests := []struct {
		name    string
		Number  string
		want    cardType
		wantErr bool
	}{
		{
			name:    "american express type check",
			Number:  "378282246310005",
			want:    AmericanExpress,
			wantErr: false,
		},
		{
			name:    "jcb type check",
			Number:  "3530111333300000",
			want:    JCB,
			wantErr: false,
		},
		{
			name:    "maestro type check",
			Number:  "6759649826438453",
			want:    Maestro,
			wantErr: false,
		},
		{
			name:    "visa type check",
			Number:  "4012888888881881",
			want:    Visa,
			wantErr: false,
		},
		{
			name:    "mastercard type check",
			Number:  "5105105105105100",
			want:    Mastercard,
			wantErr: false,
		},
		{
			name:    "Invalid card",
			Number:  "12345678901234",
			want:    -1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Card{
				Number: tt.Number,
			}
			got, err := c.GetScheme()
			if (err != nil) != tt.wantErr {
				t.Errorf("Card.GetScheme() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Card.GetScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
