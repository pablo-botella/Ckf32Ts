package main

import (
	"testing"
	"time"
	// Ajusta esta importación si el módulo está en otra ruta
)

// TestFromTime valida la conversión de time.Time a Ckf32Ts y viceversa.
func TestFromTime(t *testing.T) {
	// Tiempo de referencia
	refTime := time.Date(2025, 1, 12, 15, 30, 45, 0, time.UTC)

	// Generar CKF32TS
	ckf := Ckf32Ts.FromTime(refTime)
	ckfString := ckf.ToString()

	// Validar la representación como string
	expected := "53GJ9PQX" // Cambia por el valor esperado según tu implementación
	if ckfString != expected {
		t.Errorf("Esperado %s, pero se obtuvo %s", expected, ckfString)
	}

	// Reconstruir el tiempo desde CKF32TS
	reconstructedTime, err := ckf.ToTime()
	if err != nil {
		t.Errorf("Error al convertir CKF32TS a tiempo: %v", err)
	}

	// Validar el tiempo reconstruido
	if !reconstructedTime.Equal(refTime) {
		t.Errorf("El tiempo reconstruido no coincide. Esperado %v, obtenido %v", refTime, reconstructedTime)
	}
}

// TestToString valida la representación en string del timestamp.
func TestToString(t *testing.T) {
	// Crear un timestamp Ckf32Ts específico
	ckf := Ckf32Ts.Ckf32Ts{
		YearOffset:    5,        // 2025
		DayOfYear:     11,       // Día 12 (0-indexado)
		FractionOfDay: 14400000, // Medio día
	}

	// Convertir a string
	result := ckf.ToString()

	// Validar el resultado esperado
	expected := "53GJ8000" // Cambia este valor según tu lógica
	if result != expected {
		t.Errorf("Esperado %s, pero se obtuvo %s", expected, result)
	}
}

// TestFromString valida la creación de un Ckf32Ts desde una cadena.
func TestFromString(t *testing.T) {
	// Cadena de entrada
	input := "53GJ8000"

	// Crear desde string
	ckf, err := Ckf32Ts.FromString(input)
	if err != nil {
		t.Errorf("Error al convertir desde string: %v", err)
	}

	// Validar los valores del timestamp
	if ckf.YearOffset != 5 || ckf.DayOfYear != 11 || ckf.FractionOfDay != 14400000 {
		t.Errorf("Valores incorrectos en Ckf32Ts. Obtenido: %+v", ckf)
	}
}

// TestCaseInsensitive valida que FromString sea insensible a mayúsculas/minúsculas.
func TestCaseInsensitive(t *testing.T) {
	// Cadena de entrada en minúsculas
	input := "53gj8000"

	// Crear desde string
	ckf, err := Ckf32Ts.FromString(input)
	if err != nil {
		t.Errorf("Error al convertir desde string (case insensitive): %v", err)
	}

	// Validar los valores del timestamp
	if ckf.YearOffset != 5 || ckf.DayOfYear != 11 || ckf.FractionOfDay != 14400000 {
		t.Errorf("Valores incorrectos en Ckf32Ts (case insensitive). Obtenido: %+v", ckf)
	}
}
