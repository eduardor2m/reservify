package converters

import (
	"time"
)

func ConvertFromStringToDate(date string) (*time.Time, error) {
	formatDate := "02-01-2006"

	dateConverted, err := time.Parse(formatDate, date)

	if err != nil {
		return nil, err
	}

	return &dateConverted, nil
}