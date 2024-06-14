package datetime

import (
	"fmt"
	"time"
)

func GetDateEndOfYear(year int) (string, error) {

	dateStr := fmt.Sprintf("%d-01-02", year)

	dataTime, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		return "", err
	}

	result := time.Date(dataTime.Year(), 12, 31, 0, 0, 0, 0, time.Local).Format(time.DateOnly)

	return result, nil
}
