package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	fmt.Println("a type is :%T",a)
	f := MyFloat(11)
	fmt.Printf(" f type is : %T \n",f)
	v := Vertex{3,4}
	fmt.Printf(" v type is : %T \n",v)

	fmt.Println(f.Abs())
	fmt.Println(v.Abs())

	a = f		// a MyFloat 实现了 Abser
	fmt.Printf(" a =f ,a type is : %T, %f \n",a,a.Abs())
	a = &v		// a *Vertex 实现了 Abser
	fmt.Printf(" a =&v ,a type is : %T, %f \n",a,a.Abs())

	// a = v  //v 是一个 Vertex（而不是 *Vertex）,所以没有实现 Abser，这里 a 无法 = v
	// fmt.Println(a.Abs())

}

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X,Y float64
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}