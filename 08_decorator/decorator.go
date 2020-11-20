package decorator

import "fmt"

type Component interface {
	display()
}

type Window struct{}

func (w *Window) display() {
	fmt.Println("This is a window")
}

type TextBox struct{}

func (t *TextBox) display() {
	fmt.Println("This is a textbox")
}

type ListBox struct{}

func (l *ListBox) display() {
	fmt.Println("This is a listbox")
}

type ComponentDecorator struct {
	component Component
}

func (c *ComponentDecorator) DecorateComponent(component Component) {
	c.component = component
}

type ScrollBarDecorator struct {
	ComponentDecorator
}

func (s *ScrollBarDecorator) display() {
	s.setScrollBar()
	s.component.display()
}

func (s *ScrollBarDecorator) setScrollBar() {
	fmt.Println("add scrollbar")
}

type BlackBorderDecorator struct {
	ComponentDecorator
}

func (b *BlackBorderDecorator) display() {
	b.setBlackBorder()
	b.component.display()
}

func (b *BlackBorderDecorator) setBlackBorder() {
	fmt.Println("add blackborder")
}
