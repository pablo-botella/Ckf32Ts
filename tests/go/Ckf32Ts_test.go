package main

import (
	"fmt"
	"time"

	Ckf32Ts "github.com/pablo-botella/Ckf32Ts/Go"
)

func main() {
	now := time.Now().UTC()
	ckf := Ckf32Ts.FromTime(now)
	fmt.Println("CKF32TS generado:", ckf.ToString())

	timeFromCkf, err := ckf.ToTime()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tiempo reconstruido:", timeFromCkf)
}
