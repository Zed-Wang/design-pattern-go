package factory

import "fmt"

type Shape interface {
	draw()
}

type rectangle struct{}

func (r rectangle) draw() {
	fmt.Println("Inside Rectangle::draw() method")
}

type square struct{}

func (s square) draw() {
	fmt.Println("Inside Square::draw() method")
}

type circle struct{}

func (c circle) draw() {
	fmt.Println("Inside Circle::draw() method")
}

func NewShape(shapeType string) Shape {
	switch shapeType {
	case "rectangle":
		return rectangle{}
	case "square":
		return square{}
	case "circle":
		return circle{}
	}
	return nil
}
