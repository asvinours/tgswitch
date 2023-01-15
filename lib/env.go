package lib

import (
	"errors"
	"os"
	"strconv"
)

var errEnvVarEmpty = errors.New("getenv: environment variable empty")

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errEnvVarEmpty
	}
	return v, nil
}

func GetEnvStrWithFallback(key string, fallback string) string {
	value, err := GetEnvStr(key)
	if err != nil {
		return fallback
	}
	return value
}

func GetEnvBool(key string) (bool, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func GetEnvBoolWithFallback(key string, fallback bool) bool {
	v, err := GetEnvBool(key)
	if err != nil {
		return fallback
	}
	return v
}

func CheckEnvExist(key string) bool {
	return os.Getenv(key) != ""
}
