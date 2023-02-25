package pkg1

import "math"

func KaprekarNumber(num int64) bool {
	var numSq int64 = num * num
	var t int64 = numSq
	var c int64 = 0
	var a int64=0
	var b int64=0
	for numSq > 0 {
		c+=1
		numSq = numSq / 10
	}
	a=int64(math.Pow(10,float64(c/2)))
	b=a*10
    return (c%2==0 && (t%a+t/a)==num)||(c%2!=0 && (t%b+t/b)==num)
}
func Is_Pronic_Number(n float64) bool {
	a := 2 * math.Sqrt(n+0.25)
	b := (a + 1) * 0.5
	return b-math.Trunc(b) == 0 && b*(b-1) == n
}
func Add_digit(n int ,d int ,p int ) int{
	return n+d+p;
}