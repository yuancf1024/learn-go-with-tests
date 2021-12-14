package main

import (
	"fmt"
	"math"
	"strconv"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectangle) float64 {
	return 2*(rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// 一个奇葩的四舍五入方法
// 先+0.5，然后向下取整！
// func round(x float64) int {
// 	return int(math.Floor(x + 0.5))
// }

// 加上 0.5是为了四舍五入，想保留几位小数的话把2改掉即可。
func Decimal1(value float64) float64 {
    return math.Trunc(value*1e2+0.5) * 1e-2
}

// 下面的是先通过Sprintf保留两位小数，再转成float64.
func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func main() {
	circle := Circle{10} // radius = 10
	// got := Decimal(circle.Area())
	got := circle.Area()
	want := 314.16

	fmt.Println("*****circle******")
	// fmt.Printf("got=%v", got)
	// fmt.Println()
	// fmt.Printf("want=%v", want)
	// fmt.Println()
	fmt.Println("got= ", got) 
	// got=  314.1592653589793 说明没有保留浮点位数容易导致 got != want 出错
	fmt.Println("want=", want)

	rectangle := Rectangle{12.0, 6.0}
	got1 := rectangle.Area()
	want1 := 72.0

	fmt.Println("*****rectangle******")
	fmt.Printf("got=%v\n", got1)
	fmt.Println("got= ", got1) 
	fmt.Println()
	fmt.Printf("want=%v", want1)
}

// PS D:\gocf\src\github.com> go run "d:\gocf\src\github.com\yuancf1024\learn-go-with-tests\structs\shapes.go"
// *****circle******
// got=  314.1592653589793
// want= 314.16
// *****rectangle******
// got=72
// want=72