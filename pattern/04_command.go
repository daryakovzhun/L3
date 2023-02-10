package main

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Light структура, у которой есть 2 метода turnOn и turnOff
type Light struct {
}

func NewLight() *Light {
	return &Light{}
}

func (l *Light) turnOn() {
	fmt.Println("The light is on")
}

func (l *Light) turnOff() {
	fmt.Println("The light is off")
}

// Command интерфейс с 1 методом execute()
type Command interface {
	execute()
}

// TurnOnLightCommand структура, которая умеет включать свет
type TurnOnLightCommand struct {
	theLight Light
}

func NewTurnOnLightCommand(light Light) *TurnOnLightCommand {
	return &TurnOnLightCommand{theLight: light}
}

func (onl *TurnOnLightCommand) execute() {
	onl.theLight.turnOn()
}

// TurnOffLightCommand структура, которая умеет выключать свет
type TurnOffLightCommand struct {
	theLight Light
}

func NewTurnOffLightCommand(light Light) *TurnOffLightCommand {
	return &TurnOffLightCommand{theLight: light}
}

func (offl *TurnOffLightCommand) execute() {
	offl.theLight.turnOff()
}

// Switch (invoker) вызывает команды
type Switch struct {
	flipUpCommand, flipDownCommand Command
}

func NewSwitch(fup, fdw Command) *Switch {
	return &Switch{
		flipUpCommand:   fup,
		flipDownCommand: fdw}
}

func (s *Switch) flipUp() {
	s.flipUpCommand.execute()
}

func (s *Switch) flipDown() {
	s.flipDownCommand.execute()
}

func main() {
	l := NewLight()
	switchUp := NewTurnOnLightCommand(*l)
	switchDown := NewTurnOffLightCommand(*l)

	s := NewSwitch(switchUp, switchDown)

	s.flipUp()
	s.flipDown()
}
