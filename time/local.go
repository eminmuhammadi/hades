package db

import (
	"os"
	"time"
	_ "time/tzdata"
)

// GetLocation returns the timezone location
func GetLocation(timezone string) (*time.Location, error) {
	loc, err := time.LoadLocation(timezone)

	return loc, err
}

// Now returns the current local time
func Now() time.Time {
	loc, err := GetLocation(os.Getenv("TIMEZONE"))

	if err != nil {
		panic(err)
	}

	return time.Now().In(loc)
}
