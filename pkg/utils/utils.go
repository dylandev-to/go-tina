package utils

import (
	"os"
	"time"
)

func GetCwd() string {
	cwd, _ := os.Getwd()
	return cwd
}

func CompareTimeNow(dateStr string) (*float64, error) {
	layout := "2006-01-02 15:04:05"

	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	// Calculate the difference in minutes
	duration := now.Sub(parsedDate)
	minutes := duration.Minutes()

	return &minutes, nil
}
