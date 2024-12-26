package utils

import "os"

func GetCwd() string {
	cwd, _ := os.Getwd()
	return cwd
}
