package main

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("2100000000000000000000", 10)
	b.SetString("1048576000000000000000", 10)

	z := new(big.Int)
	z.Mul(a, b)
	fmt.Println(z.String()) //2202009600000000000000000000000000000000000 

	r := new(big.Rat)
	r.SetFrac(a, b)
	fmt.Println(r.FloatString(20)) // 2.00271606445312500000

	z = new(big.Int)
	z.Add(a, b)
	fmt.Println(z.String()) // 3148576000000000000000

	z = new(big.Int)
	z.Sub(a, b)
	fmt.Println(z.String()) // 1051424000000000000000
}
