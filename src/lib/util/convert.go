package util

import (
	"strconv"
	"strings"
	"time"
)

func StrToInt(str string) int64 {
	if str != "NULL" {
		i, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return i
		} else {
			panic(err)
		}
	} else {
		return 0
	}
}

func StrToDate(str string) time.Time {
	if str != "NULL" {
		data, err := time.Parse("2006-01-02", str)

		if err == nil {
			return data
		} else {
			panic(err)
		}

	} else {
		data, err := time.Parse("2006-01-02", "1900-01-01")
		if err == nil {
			return data
		} else {
			panic(err)
		}
	}
}

func StrToFloat(str string) float64 {
	if str != "NULL" {
		str := strings.Replace(str, ",", ".", -1)
		value, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return value
		} else {
			panic(err)
		}
	} else {
		return 0
	}
}