package main

import (
	"demo03/ch4/Struct/Circle"
	"demo03/ch4/Struct/Point"
	"demo03/ch4/Struct/Wheel"
	"fmt"
)

func main() {

	var wheel Wheel.Wheel

	wheel = Wheel.Wheel{
		Circle: Circle.Circle{
			Point: Point.Point{
				X: 1,
				Y: 2,
			},
			Radius: 3,
		},
		Spokes: 4,
	}

	fmt.Println(wheel)
}

