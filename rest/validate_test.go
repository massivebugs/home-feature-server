package rest

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestIsValidDateTime(t *testing.T) {
	type s struct {
		DateTime string `validate:"_iso8601"`
	}

	validStruct := s{
		DateTime: "2024-11-18T00:00:30.465Z",
	}

	invalidStruct := s{
		DateTime: "2024-11-18 00:00:30",
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
			name: "success when value is correct datetime format",
			args: args{
				s: validStruct,
			},
			wantErr: false,
		},
		{
			name: "error when value is not the correct datetime format",
			args: args{
				s: invalidStruct,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validator.New()
			v.RegisterValidation("_iso8601", IsValidDateTime)

			if err := v.Struct(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("IsValidDateTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
