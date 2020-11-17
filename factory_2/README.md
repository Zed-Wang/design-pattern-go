## Go语言设计模式02——工厂模式

工厂模式（Factory Pattern）是一种创建型模式，它提供了一种创建对象的最佳方式。在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。

### 简单工厂模式

```go
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
```

创建一个 Shape 接口和一些实现 Shape 接口的实体类，并通过 ShapeFactory 来获取 Shape。在Go语言中，可以直接使用 NewShape 函数来获取

### 工厂模式

```go
package factory

type Shape interface {
	draw()
}

type rectangle struct{}

type square struct{}

type circle struct{}

type ShapeFactory interface {
	CreateShape() Shape
}

type rectangleFactory struct{}

func (r rectangleFactory) CreateShape() Shape {
	return rectangle{}
}

type squareFactory struct{}

func (c squareFactory) CreateShape() Shape {
	return square{}
}

type circleFactory struct{}

func (c circleFactory) CreateShape() Shape {
	return circle{}
}

func NewShapeFactory(shapeFactoryType string) ShapeFactory {
	switch shapeFactoryType {
	case "rectangle":
		return rectangleFactory{}
	case "square":
		return squareFactory{}
	case "circle":
		return circleFactory{}
	}
	return nil
}
```

工厂模式比简单工厂模式复杂一点，可以理解为简单工厂外面再套一个工厂。创建 ShapeFactory 接口和实现 ShapeFactory 的实体类，这些实体类又是各种 Shape 的工厂类，这些细分的工厂类统一再用一个工厂类来调用。在Go语言中，我们直接使用 NewShapeFactory 来获取细分的工厂类