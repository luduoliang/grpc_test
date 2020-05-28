package config

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	//当前文件路劲
	execPath, _ := os.Getwd()
	_ = godotenv.Load(execPath + "/.env")
	_ = os.Setenv("APP_PATH", execPath)
}

func Get(name string) string {
	return os.Getenv(name)
}

func Default(name, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}
	return v
}

func Set(name, value string) error {
	return os.Setenv(name, value)
}

func Env() string {
	return Get("APP_ENV")
}
