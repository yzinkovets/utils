package env

import (
	"log"
	"os"
)

type iLogger interface {
	Fatalf(format string, v ...interface{})
}

var logger iLogger

// Set default logger
func init() {
	logger = log.New(os.Stdout, "env: ", log.LstdFlags)
}

func SetLogger(l iLogger) {
	logger = l
}

func Get(name string) string {
	return get(name, false)
}
func Must(name string) string {
	return get(name, true)
}

func GetDef(name string, def string) string {
	env, ok := os.LookupEnv(name)
	if !ok {
		return def
	}
	return env
}

func GetBoolDef(name string, def bool) bool {
	env, ok := os.LookupEnv(name)
	if !ok {
		return def
	}
	return env == "true" || env == "1"
}

func get(name string, required bool) string {
	env, ok := os.LookupEnv(name)
	if !ok && required {
		logger.Fatalf("Please set %s environment variable\n", name)
	}
	return env
}
