package Ckf32Ts

import (
	"errors"
	"time"
)

// Base 32 alfabeto
var base32Alphabet = []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K',
	'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z'}

// Mapa para decodificar base 32 (case-insensitive)
var base32Map = func() map[rune]int {
	m := make(map[rune]int)
	for i, char := range base32Alphabet {
		m[char] = i
		m[char+('a'-'A')] = i // Soporte para minúsculas
	}
	return m
}()

// Ckf32Ts estructura para el timestamp
type Ckf32Ts struct {
	YearOffset    int
	Month         int
	Day           int
	FractionOfDay uint64
}

func FromTime(t time.Time) Ckf32Ts {
	yearOffset := t.Year() - 2020
	month := int(t.Month())
	day := t.Day()

	hours := t.Hour()
	minutes := t.Minute()
	seconds := t.Second()

	secondsToday := (hours * 3600) + (minutes * 60) + seconds
	millisecondsToday := secondsToday * 1000
	fractionOfDay := (uint64(millisecondsToday) * 28800000) / 86400000

	return Ckf32Ts{
		YearOffset:    yearOffset,
		Month:         month,
		Day:           day,
		FractionOfDay: fractionOfDay,
	}
}

// ToTime convierte Ckf32Ts a time.Time
func (ckf Ckf32Ts) ToTime() (time.Time, error) {
	if ckf.YearOffset < 0 || ckf.YearOffset > 31 {
		return time.Time{}, errors.New("año fuera de rango")
	}
	if ckf.Month < 1 || ckf.Month > 12 {
		return time.Time{}, errors.New("mes fuera de rango")
	}
	if ckf.Day < 1 || ckf.Day > 31 {
		return time.Time{}, errors.New("día fuera de rango")
	}

	year := 2020 + ckf.YearOffset
	secondsInDay := (ckf.FractionOfDay * 86400000) / 28800000
	hours := secondsInDay / 3600
	minutes := (secondsInDay % 3600) / 60
	seconds := secondsInDay % 60

	loc := time.UTC
	return time.Date(year, time.Month(ckf.Month), ckf.Day, int(hours), int(minutes), int(seconds), 0, loc), nil
}

// ToString convierte un Ckf32Ts a su representación como string
func (ckf Ckf32Ts) ToString() string {
	return string(base32Alphabet[ckf.YearOffset]) +
		toBase32(uint64(ckf.Month), 1) +
		toBase32(uint64(ckf.Day), 1) +
		toBase32(ckf.FractionOfDay, 5)
}

// Función auxiliar para convertir a base 32
func toBase32(value uint64, length int) string {
	result := make([]rune, length)
	for i := length - 1; i >= 0; i-- {
		result[i] = base32Alphabet[value%32]
		value /= 32
	}
	return string(result)
}

// FromString convierte una cadena CKF32TS a la estructura Ckf32Ts
func FromString(input string) (Ckf32Ts, error) {
	if len(input) != 8 {
		return Ckf32Ts{}, errors.New("la cadena CKF32TS debe tener exactamente 8 caracteres")
	}

	// Decodificar el año (primer carácter)
	yearOffset, exists := base32Map[rune(input[0])]
	if !exists {
		return Ckf32Ts{}, errors.New("carácter inválido en el año")
	}

	// Decodificar el mes (carácter 2)
	month, err := fromBase32(input[1:2])
	if err != nil {
		return Ckf32Ts{}, errors.New("error al decodificar el mes")
	}

	// Decodificar el día (carácter 3)
	day, err := fromBase32(input[2:3])
	if err != nil {
		return Ckf32Ts{}, errors.New("error al decodificar el día")
	}

	// Decodificar la fracción del día (caracteres 4 al 8)
	fractionOfDay, err := fromBase32(input[3:8])
	if err != nil {
		return Ckf32Ts{}, errors.New("error al decodificar la fracción del día")
	}

	// Construir la estructura
	return Ckf32Ts{
		YearOffset:    yearOffset,
		Month:         int(month),
		Day:           int(day),
		FractionOfDay: fractionOfDay,
	}, nil
}

// Función auxiliar para convertir cadenas base 32 a enteros
func fromBase32(input string) (uint64, error) {
	var value uint64
	for _, char := range input {
		index, exists := base32Map[char]
		if !exists {
			return 0, errors.New("carácter inválido en base 32")
		}
		value = value*32 + uint64(index)
	}
	return value, nil
}

// TimeToCkf32TsString convierte un objeto time.Time a una cadena CKF32TS
func TimeToCkf32TsString(t time.Time) string {
	ckf := FromTime(t)
	return ckf.ToString()
}

// Ckf32TsStringToTime convierte una cadena CKF32TS a un objeto time.Time
func Ckf32TsStringToTime(input string) (time.Time, error) {
	ckf, err := FromString(input)
	if err != nil {
		return time.Time{}, err
	}
	return ckf.ToTime()
}
