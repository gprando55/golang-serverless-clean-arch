package date

import (
	"strconv"
	"time"
)

type date struct {
	current time.Time
}

func Now() *date {
	return &date{current: time.Now()}
}

func OfPattern(toParse, pattern string) (*date, error) {
	dateParsed, err := time.Parse(pattern, toParse)

	if err != nil {
		return nil, err
	}

	return &date{current: dateParsed}, nil
}
func (d date) DynamoFormat() int {
	formatted, _ := strconv.Atoi(d.current.Format("20060102150405"))

	return formatted
}
