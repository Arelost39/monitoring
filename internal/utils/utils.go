package utils

import (
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv загружает переменные окружения из файла .env.
// Вызывается один раз при инициализации приложения.
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		Log.Errorf("Ошибка загрузки .env файла: %v", err)
		return err
	}
	return nil
}

// GetEnv возвращает значение переменной окружения для указанного ключа.
func GetEnv(key string) string {
	return os.Getenv(key)
}
