package factory

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
