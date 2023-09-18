package main

// Реализовать паттерн «адаптер» на любом примере.

import (
    "fmt"
)

type CelsiusTemp interface {
    GetTempCelsius() float64
}

type FahrenheitTemp struct {
    Temp float64
}

type TempAdapter struct {
    FahrenheitTemp
}

func (f TempAdapter) GetTempCelsius() float64 {
    return (f.Temp - 32) * 5 / 9
}

func main() {
    fahrenheitTemp := FahrenheitTemp{Temp: 98.6}
    adapter := TempAdapter{fahrenheitTemp}

    fmt.Printf("Fahrenheit: %.2f\n", fahrenheitTemp.Temp) // Fahrenheit: 98.60
    fmt.Printf("Celsius: %.2f\n", adapter.GetTempCelsius()) // Celsius: 37.00
}