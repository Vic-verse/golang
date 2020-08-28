package main

import (
	"fmt"
)

func main() {
	fmt.Print("\n\t-----Enter all requirements :----- ")
	var a, v0, s0, t float64
	fmt.Print("\n\n\tEnter the acceleration: ")
	fmt.Scanln(&a)
	fmt.Print("\n\tEnter the initial velocity: ")
	fmt.Scanln(&v0)
	fmt.Print("\n\tEnter the initial displacement: ")
	fmt.Scanln(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	fmt.Print("\n\tEnter the value of time t: ")
	fmt.Scanln(&t)
	fmt.Print("\n\tDisplacement:", fn(t))
	fmt.Print("\n\n\t----- Program closed **THANKU** and **HAVE A NICE DAY** :P -----")
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return s0 + v0*t + 0.5*a*t*t
	}
}
