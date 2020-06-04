package main;

import (
	"fmt"
)

type Animal struct {
	food, move, sound string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.move)
}

func (a Animal) Speak() {
	fmt.Println(a.sound)
}

func main() {
	m := map[string]Animal{
		"cow" : Animal{"grass","walk","moo"},
		"bird" : Animal{"worms","fly","peep"},
		"snake" : Animal{"mice","slither","hsss"},
	}
	for{
		fmt.Print(">")
		animal:=""
		action:=""
		fmt.Scan(&animal,&action)
		if action=="eat"{
			m[animal].Eat()
		} else if action=="move"{
			m[animal].Move()
		} else if action=="speak"{
			m[animal].Speak()
		}
	}
}