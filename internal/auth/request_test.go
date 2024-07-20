package auth

import "testing"

func TestUserAuthRequestDTO_Validate(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Username: "foo",
				Password: "bar@1234",
			},
			wantErr: false,
		},
		{
			name: "validation failed because no username",
			fields: fields{
				Password: "bar@1234",
			},
			wantErr: true,
		},
		{
			name: "validation failed because username is too short",
			fields: fields{
				Username: "fo",
				Password: "bar@1234",
			},
			wantErr: true,
		},
		{
			name: "validation failed because username is too long",
			fields: fields{
				Username: "foooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo",
				Password: "bar@1234",
			},
			wantErr: true,
		},
		{
			name: "validation failed because username is invalid",
			fields: fields{
				Username: "fo@",
				Password: "bar@1234",
			},
			wantErr: true,
		},
		{
			name: "validation failed because no password",
			fields: fields{
				Username: "foo",
			},
			wantErr: true,
		},
		{
			name: "validation failed because password is invalid",
			fields: fields{
				Username: "foo",
				Password: "bar",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &UserAuthRequestDTO{
				Username: tt.fields.Username,
				Password: tt.fields.Password,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("UserAuthRequestDTO.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
