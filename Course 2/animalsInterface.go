package main

import "fmt"

// Animal interface.
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow animal definition
type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird animal definition
type Bird struct{}

func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}

// Snake animal definition
type Snake struct{}

func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

var animalsMap map[string]Animal

func createAnimal(name, request string) {
	var animal Animal
	switch request {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Println("Invalid animal type.")
		return
	}
	fmt.Println("Created it!")
	animalsMap[name] = animal
}

func queryAnimals(name, request string) {
	animal, exist := animalsMap[name]
	if !exist {
		fmt.Printf("An animal with name %v doesn't exist.\n", name)
		return
	}
	switch request {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid animal type.")
	}
}

func main() {
	animalsMap = make(map[string]Animal)
	var action, name, request string
	for {
		fmt.Print("> ")
		fmt.Scan(&action, &name, &request)

		switch action {
		case "newanimal":
			createAnimal(name, request)
		case "query":
			queryAnimals(name, request)
		default:
			fmt.Println("Invalid action.")
		}
	}
}
