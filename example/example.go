package main

import "fmt"

type Position struct {
	X, Y, Z uint64
}

type Weapon struct {
	Name   string
	Damage uint16
}

type Player struct {
	Name    string
	Health  uint16
	Weapons []Weapon

	Position
}

func main() {
	player1 := Player{
		Name:    "Jane",
		Health:  100,
		Weapons: []Weapon{{Name: "Sword", Damage: 10}},
		Position: Position{
			X: 1,
			Y: 2,
			Z: 3,
		},
	}

	fmt.Printf("player1: %v\n", player1)

	size := player1.SizeGOBE()
	fmt.Printf("Size of player1: %d bytes\n", size)

	buffer := make([]byte, size)
	player1.MarshalGOBE(buffer)

	fmt.Printf("buffer: %v\n", buffer)

	fmt.Println("Unmarshal player2 from buffer")
	player2 := Player{}
	_, ok := player2.UnmarshalGOBE(buffer)
	if !ok {
		fmt.Println("Unmarshal failed")
	}
	fmt.Printf("player2: %v\n", player2)
}
