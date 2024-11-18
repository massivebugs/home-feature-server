package rest

import (
	"testing"
	"time"

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
			v.RegisterValidation("_iso8601", IsValidDateTime(time.RFC3339Nano))

			if err := v.Struct(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("IsValidDateTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsValidDateTimeFormat(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success valid time format returns no errors",
			args: args{
				v: "2006-01-02T15:04:05.999999999Z07:00",
			},
			wantErr: false,
		},
		{
			name: "success valid time format with no ms precision returns no errors",
			args: args{
				v: "2006-01-02 15:04:05",
			},
			wantErr: false,
		},
		{
			name: "error invalid time format",
			args: args{
				v: "foo",
			},
			wantErr: true,
		},
		{
			name:    "error empty string",
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsValidDateTimeFormat(tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("IsValidDateTimeFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
