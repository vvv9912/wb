package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.

Паттерн Адаптер (Adapter) предназначен для преобразования интерфейса одного класса в интерфейс другого.
Благодаря реализации данного паттерна мы можем использовать вместе классы с несовместимыми интерфейсами.
*/
//Задача имеется порт USB и COM, необходимо принять сообщение

// интерфейс описывающий имеющий парсер  usb порта
type messeger interface {
	usbParser()
}
type usb struct {
}

func (u *usb) usbParser() {
	fmt.Println("Информация обработалась с usb порта")
}

// опишем ф-ции для com
type com struct {
}

func (c *com) comParser() {
	fmt.Println("Информация обработалась с com порта")
}

// адаптер
type comAdapter struct {
	comPort *com
}

func (c comAdapter) usbParser() {
	c.comPort.comParser()
}

// Обработка сообщения
type message struct {
}

// сама ф-ия
func (m *message) parser(mes messeger) {
	mes.usbParser()
}

func main() {
	com := &com{}
	usb := &usb{}
	comAdapterr := &comAdapter{comPort: com}
	//
	usb.usbParser()
	comAdapterr.usbParser()
	//
	messagee := &message{}
	messagee.parser(comAdapterr)
	messagee.parser(usb)
}
