package adapter

import "fmt"

type CarController interface {
	phonate()
	twinkle()
}

type PoliceSound struct{}

func (p *PoliceSound) alarmSound() {
	fmt.Println("发出警笛声音！")
}

type PoliceLamp struct{}

func (p *PoliceLamp) alarmLamp() {
	fmt.Println("呈现警灯闪烁！")
}

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
