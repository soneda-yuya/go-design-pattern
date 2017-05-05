package main

import (
	"fmt"
)

type Athlete struct{}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

type SwimmerA struct{
	MyAthlete Athlete
	MySwim func()
}

type Animal struct{}

func (r *Animal)Eat() {
	println("Eating")
}

type Shark struct{
	Animal
	Swim func()
}

func Swim(){
	fmt.Println("Swimming!")
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim(){
  println("Swimming!!!")
}

type SwimmerB struct{
	Trainer
	Swimmer
}

func main(){
	// Pattern 1
	swimmer := SwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// Pattern 2
	localSwim := Swim

	swimmer = SwimmerA{
	  MySwim: localSwim,
	}

	swimmer.MyAthlete.Train()
	swimmer.MySwim()

	// Pattern 3
	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	// Pattern 4
	swimmer2 := SwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer2.Train()
	swimmer2.Swim()
}