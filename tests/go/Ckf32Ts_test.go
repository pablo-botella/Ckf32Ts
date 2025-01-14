package main

import (
	"testing"
	"time"

	Ckf32Ts "github.com/pablo-botella/Ckf32Ts/go"
)

func TestFromTime(t *testing.T) {
	refTime := time.Date(2025, 1, 12, 15, 30, 45, 0, time.UTC)

	// Generar CKF32TS
	ckf := Ckf32Ts.FromTime(refTime)

	// Validar la representación como string
	expected := "51CHR2PR" // Nueva representación esperada según el cálculo y formato ajustado
	t.Logf("Valor obtenido: %s", ckf.ToString())
	if ckf.ToString() != expected {
		t.Errorf("Esperado %s, pero se obtuvo %s", expected, ckf.ToString())
	}
}
