## Go语言设计模式08——装饰模式

装饰模式是一种结构型模式，可以看做是替代继承的技术。它通过一种无须定义子类的方式给对象动态增加职责，使用对象之前的关联取代类之间的继承关系

### 实例

有个公司开发了一套图形界面构件库,如 Window，TextBox，ListBox，都有基本的显示功能

```go
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
```

但是有时用户需要定制一些特殊的显示效果，如带滚动条的窗体、带黑色边框的文本框、既带滚动条又带黑色边框的列表框等

因此创建抽象装饰器 ComponentDecorator

```go
type ComponentDecorator struct {
	component Component
}

func (c *ComponentDecorator) DecorateComponent(component Component) {
	c.component = component
}
```

和两个具体装饰器

```go
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
```

调用

```go
func main() {
	component := new(Window)
	componentSB := new(ScrollBarDecorator)
	componentBB := new(BlackBorderDecorator)
	componentSB.DecorateComponent(component)
	componentBB.DecorateComponent(componentSB)
	componentBB.display()
}
```

结果

```
add blackborder
add scrollbar
This is a window
```