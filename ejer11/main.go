package main

import "fmt"

type LivingBeing interface {
	isAlive() bool
}

type Human interface {
	breathe()
	think()
	eat()
	sex() string
}

type Animal interface {
	breathe()
	eat()
	isCarnivorous() bool
}

type Vegetal interface {
	Classification() string
}

// Género Humano
type Man struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	isMan     bool
	live      bool
}

type Woman struct {
	Man
}

func (m *Man) isAlive() bool {
	return m.live
}
func (m *Man) breathe() {
	m.breathing = true
}
func (m *Man) eat() {
	m.eating = true
}
func (m *Man) think() {
	m.thinking = true
}
func (m *Man) sex() string {
	if m.isMan {
		return "Man"

	} else {
		return "Woman"

	}
}

func HumanBreathing(hu Human) {
	hu.breathe()
	fmt.Printf("I'm a %s and i breathing \n", hu.sex())
}

/*----------------------------------------------------*/
// Género Animal
type Dog struct {
	breathing   bool
	eating      bool
	carnivorous bool
	live        bool
}

func (d *Dog) breathe() {
	d.breathing = true
}
func (d *Dog) eat() {
	d.eating = true
}
func (d *Dog) isCarnivorous() bool {
	return d.carnivorous
}
func (d *Dog) isAlive() bool {
	return d.live
}

func AnimalBreathing(a Animal) {
	a.breathe()
	fmt.Printf("I'm a Animal and i breathing \n")
}

func AnimalCarnivorous(a Animal) int {
	if a.isCarnivorous() {
		return 1
	}
	return 0
}

func ImAlive(lb LivingBeing) {
	if lb.isAlive() {
		fmt.Print("I'm  a living being\n")
	} else {
		fmt.Print("I'am not a living being\n")
	}
}

/*----------------------------------------------------*/
// Género Animal

func main() {
	Pedro := new(Man)
	Pedro.isMan = true
	Pedro.live = true
	HumanBreathing(Pedro)

	Paola := new(Woman)
	Paola.live = true
	HumanBreathing(Paola)

	totalCarnivorous := 0
	Dogo := new(Dog)
	Dogo.carnivorous = true
	Dogo.live = true
	AnimalBreathing(Dogo)
	totalCarnivorous += AnimalCarnivorous(Dogo)
	fmt.Printf("Total Carnivorous %d\n", totalCarnivorous)

	ImAlive(Pedro)
	ImAlive(Paola)
	ImAlive(Dogo)
}
