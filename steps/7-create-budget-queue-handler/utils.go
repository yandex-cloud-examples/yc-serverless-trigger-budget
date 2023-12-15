package main

import (
	"log"
	"os"
	"strconv"
)

func requireEnvInt64(envVarName string) int64 {
	res, err := strconv.ParseInt(os.Getenv(envVarName), 10, 64)
	if err != nil {
		log.Fatalf("could not parse '%s' env var as int64: %v", envVarName, err)
	}
	return res
}
