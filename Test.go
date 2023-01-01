package main

import (
	"awesomeProject/pkg1"
	"awesomeProject/pkg2"
)

func main() {
	println(pkg1.KaprekarNumber(45))
	println(pkg1.Is_Pronic_Number(72))
	pkg2.Start()
	pkg1.Add_digit(5,9,10)
	pkg1.Heaps_permutation([]int{2, 3, 5,8},4)
}
