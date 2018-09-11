package main

import "fmt"

func sqrt(x float64) float64 {
	const a  =  0.000001
	z := float64(1)
	k := float64(0)
	for ;; z = z - (z*z-x)/(2*z) {
		if z-k <= a && z-k >= -a {
			return z
		}
		k=z
	}
}
func main()  {
	fmt.Println(sqrt(2))
}
