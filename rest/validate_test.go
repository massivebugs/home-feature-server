package rest

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ValidatableStruct struct {
	SomeRequiredField string
}

func (s *ValidatableStruct) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(
			&s.SomeRequiredField,
			validation.Required,
		),
	)
}

type StructWithoutValidation struct{}

func TestRequestValidator_Validate(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "validation success",
			args: args{
				i: &ValidatableStruct{
					SomeRequiredField: "foo",
				},
			},
			wantErr: false,
		},
		{
			name: "validation error",
			args: args{
				i: &ValidatableStruct{},
			},
			wantErr: true,
		},
		{
			name: "error because struct does not implement validate.Validatable",
			args: args{
				i: &StructWithoutValidation{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RequestValidator{}
			if err := r.Validate(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("RequestValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
