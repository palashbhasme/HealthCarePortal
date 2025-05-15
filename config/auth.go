package config

type AuthConfig struct {
	SecretKey string
}

func NewAuthConfig(secretKey string) *AuthConfig {
	return &AuthConfig{
		SecretKey: secretKey,
	}
}
