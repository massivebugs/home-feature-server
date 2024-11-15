package auth

// func TestIsValidPassword(t *testing.T) {
// 	type args struct {
// 		minLength int
// 		maxLength int
// 		value     string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "password is valid",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 72,
// 				value:     "abcdef1@",
// 			},
// 		},
// 		{
// 			name: "password is too short",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 72,
// 				value:     "a1@",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "password is too long",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 10,
// 				value:     "abcdef1@@@@@",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "password does not include number",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 72,
// 				value:     "aaaaa@@@@@",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "password does not include character",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 72,
// 				value:     "aaaaa11111",
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "password does not include letter",
// 			args: args{
// 				minLength: 8,
// 				maxLength: 72,
// 				value:     "11111@@@@@",
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			validateFunc := isValidPassword(tt.args.minLength, tt.args.maxLength)
// 			if err := validateFunc(tt.args.value); (err != nil) != tt.wantErr {
// 				t.Errorf("IsValidPassword() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
