package Ckf32Ts

import (
	"errors"
	"time"
)

// Base 32 alfabeto
var base32Alphabet = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K',
	'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}

// Ckf32Ts estructura para el timestamp
type Ckf32Ts struct {
	YearOffset    int
	DayOfYear     int
	FractionOfDay uint64
}

// FromTime convierte time.Time a Ckf32Ts
func FromTime(t time.Time) Ckf32Ts {
	yearOffset := t.Year() - 2020
	dayOfYear := t.YearDay() - 1
	secondsToday := t.Hour()*3600 + t.Minute()*60 + t.Second()
	millisecondsToday := secondsToday * 1000
	fractionOfDay := (uint64(millisecondsToday) * 28800000) / 86400000

	return Ckf32Ts{
		YearOffset:    yearOffset,
		DayOfYear:     dayOfYear,
		FractionOfDay: fractionOfDay,
	}
}

// ToTime convierte Ckf32Ts a time.Time
func (ckf Ckf32Ts) ToTime() (time.Time, error) {
	if ckf.YearOffset < 0 || ckf.YearOffset > 31 {
		return time.Time{}, errors.New("año fuera de rango")
	}
	if ckf.DayOfYear < 0 || ckf.DayOfYear > 366 {
		return time.Time{}, errors.New("día del año fuera de rango")
	}

	year := 2020 + ckf.YearOffset
	secondsInDay := (ckf.FractionOfDay * 86400000) / 28800000
	hours := secondsInDay / 3600
	minutes := (secondsInDay % 3600) / 60
	seconds := secondsInDay % 60

	loc := time.UTC
	return time.Date(year, 1, 1, int(hours), int(minutes), int(seconds), 0, loc).AddDate(0, 0, ckf.DayOfYear), nil
}
