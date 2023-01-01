package pkg2

import (
	"fmt"
	"math"
)

type Vertex struct {
	X int
	Y int
}

func Start() {
	list := [3]float64{50, 15, 60}
	fmt.Println(findHcfAndLcm(list))
	x := Vertex{X: 5}
	println(x.Y)
	example()
}
func findHcfAndLcm(arr [3]float64) (float64, float64) {
	l := arr[0]
	h := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		h = hcf(h, arr[i+1])
		l = lcm(l, arr[i+1])
	}
	return l, h
}
func hcf(a float64, b float64) float64 {
	max := math.Max(a, b)
	min := math.Min(a, b)
	tem := 0.0
	for math.Mod(max, min) != 0 {
		tem = min
		min = math.Mod(max, min)
		max = tem
	}
	return min
}
func lcm(a float64, b float64) float64 {
	return (a * b) / hcf(a, b)
}
func example() {
	i, j := 42, 2246846847
	s := "Hello"
	p := &i // point to i
	q := &s
	fmt.Println(p, &j, q) // read i through the pointer

	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	println(p)
	*p = *p + 36/2 + 3 // divide j through the pointer
	fmt.Println(j)     // see the new value of j
}
