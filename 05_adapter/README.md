## Go语言设计模式05——适配器模式

适配器模式是一种结构性模式。如果系统中存在不兼容的接口，可以通过引入一个适配器来使原本因为接口不兼容而不能一起工作的两个类能够协同工作

通俗来讲，我理解的适配器模式就是避免修改函数名，通过外面套一个壳子来实现接口

### 实例

有个公司开发了一款玩具汽车，具有灯光闪烁和发出声音的功能

```go
type CarController interface {
	phonate()
	twinkle()
}
```

同时，公司在以前的产品中已经实现了**警灯闪烁**和**警笛音效**的功能

```go
type PoliceSound struct{}

func (p *PoliceSound) alarmSound() {
	fmt.Println("发出警笛声音！")
}

type PoliceLamp struct{}

func (p *PoliceLamp) alarmLamp() {
	fmt.Println("呈现警灯闪烁！")
}
```

因此，公司想要把这两个功能替换到玩具汽车上，所以编写适配器

```go
type PoliceCarAdapter struct {
	sound PoliceSound
	lamp  PoliceLamp
}

// 发出警笛声音
func (p *PoliceCarAdapter) phonate() {
	p.sound.alarmSound()
}

// 呈现警灯闪烁
func (p *PoliceCarAdapter) twinkle() {
	p.lamp.alarmLamp()
}
```

调用

```go
func main() {
	sound := PoliceSound{}
	lamp := PoliceLamp{}
	car := PoliceCarAdapter{sound, lamp}
	car.phonate()
	car.twinkle()
}
```

