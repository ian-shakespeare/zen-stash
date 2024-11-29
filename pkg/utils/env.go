package utils

import (
	"fmt"
	"os"
)

func RequiredEnv(name string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		msg := fmt.Sprintf("missing required env %s", name)
		panic(msg)
	}
	return value
}

func FallbackEnv(name, fallback string) string {
	value, exists := os.LookupEnv(name)
	if !exists {
		return fallback
	}
	return value
}
