package converters

import (
	"fmt"
	"strconv"
	"time"
)

func ConverterFromStringToTime(date string) (time.Time, error) {
	dateFormat := "02/01/2006"

	dateTime, err := time.Parse(dateFormat, date)

	if err != nil {
		return time.Time{}, fmt.Errorf("falha ao converter data: %v", err)
	}

	return dateTime, nil
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}