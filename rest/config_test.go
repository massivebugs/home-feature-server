package rest

import (
	"testing"
)

func TestConfig_Load(t *testing.T) {
	tests := []struct {
		name       string
		setEnvFunc func(t *testing.T)
		wantErr    bool
	}{
		{
			name: "Success all fields are present as environment variables",
			setEnvFunc: func(t *testing.T) {
				t.Setenv("ENVIRONMENT", "local")
				t.Setenv("API_PORT", "1323")
				t.Setenv("ALLOWED_ORIGINS", "https://localhost:5173")
				t.Setenv("TLS_CERTIFICATE", "devcerts/localhost.pem")
				t.Setenv("TLS_KEY", "devcerts/localhost-key.pem")
				t.Setenv("AUTH_JWT_COOKIE_NAME", "auth_token")
				t.Setenv("AUTH_JWT_SIGNING_METHOD_HMAC", "HS256")
				t.Setenv("AUTH_JWT_SECRET", "test_secret")
				t.Setenv("AUTH_JWT_EXPIRE_SECONDS", "900")
				t.Setenv("REFRESH_JWT_SIGNING_METHOD_HMAC", "HS256")
				t.Setenv("REFRESH_JWT_SECRET", "test_secret_2")
				t.Setenv("REFRESH_JWT_EXPIRE_SECONDS", "86400")
				t.Setenv("MYSQL_HOST", "db")
				t.Setenv("MYSQL_PORT", "3306")
				t.Setenv("MYSQL_DATABASE", "home_feature_server")
				t.Setenv("MYSQL_USER", "home_feature_server")
				t.Setenv("MYSQL_PASSWORD", "password")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setEnvFunc(t)
			c := NewConfig()
			if err := c.Load(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
