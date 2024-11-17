package cashbunny

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestIsValidCurrency(t *testing.T) {
	type s struct {
		Currency string `validate:"_cashbunny_currency"`
	}

	validStruct := s{
		Currency: "CAD",
	}

	invalidStruct := s{
		Currency: "FOO",
	}

	type args struct {
		s s
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success when value is supported currency",
			args: args{
				s: validStruct,
			},
			wantErr: false,
		},
		{
			name: "error when value is not supported currency",
			args: args{
				s: invalidStruct,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			v.RegisterValidation("_cashbunny_currency", IsValidCurrency)

			if err := v.Struct(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("IsValidCurrency() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
