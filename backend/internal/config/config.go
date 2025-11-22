package config

func Get(key string) (string, error) {
	return "postgresql://avito_user:avito_password@localhost:5432/avito_db?sslmode=disable", nil
	// if val := os.Getenv(key); val != "" {
	// 	return val, nil
	// }
	// return "", fmt.Errorf("have not acces to env. {config.go}")
}
