package main

import (
	"fmt"
)

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}


type Memento struct {
	memento Command
}

type originator struct {
	Command Command
}

func (o *originator) NewMemento() Memento {
	return Memento{memento: o.Command}
}

func (o *originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.memento
}

type careTaker struct {
	mementoList []Memento
}

func (c *careTaker) Add(m Memento) {
	c.mementoList = append(c.mementoList, m)
}

func (c *careTaker) Pop() Memento {
  if len(c.mementoList) > 0 {
		tempMemento := c.mementoList[len(c.mementoList)-1]
		c.mementoList = c.mementoList[0:len(c.mementoList)-1]
		return tempMemento
  }

  return Memento{}
}

type MementoFacade struct {
  originator originator
  careTaker  careTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.originator.Command = s
	m.careTaker.Add(m.originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings(i int) Command {
	m.originator.ExtractAndStoreCommand(m.careTaker.Pop())
	return m.originator.Command
}

func main(){
	m := MementoFacade{}

	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))
	assertAndPrint(m.RestoreSettings(0))
	assertAndPrint(m.RestoreSettings(1))
}

func assertAndPrint(c Command){
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}