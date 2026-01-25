package dp

import (
	"os"
	"time"
)

func LoadAppTimeZone() (*time.Location, error) {
	timezone := os.Getenv("TIME_ZONE")
	if timezone == "" {
		timezone = "Europe/Berlin"
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		return nil, err
	}
	return location, nil
}
