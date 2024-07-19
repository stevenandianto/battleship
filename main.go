package main

import (
	"battleship/model"
	"battleship/usecase"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Split(r rune) bool {
	return r == ':' || r == ' '
}

func main() {
	var inputs = usecase.ScanFile("./input.txt")
	fmt.Println(inputs[2])

	// initialize battlefield
	size, err := strconv.Atoi(inputs[0])
	if err != nil {
		log.Fatal(err)
	}

	// initialize matrix
	field := make([][]model.Ship, size)
	for i := 0; i < size; i++ {
		field[i] = make([]model.Ship, size)
	}

	// initialize player

	player1 := model.Player{PlayerID: 1}
	player2 := model.Player{PlayerID: 2}

	coordinatesShip1 := strings.FieldsFunc(inputs[2], Split)
	for _, v := range coordinatesShip1 {
		pointers := strings.Split(v, ",")
		pointerX, err := strconv.Atoi(pointers[0])
		if err != nil {
			log.Fatal(err)
		}

		pointerY, err := strconv.Atoi(pointers[1])
		if err != nil {
			log.Fatal(err)
		}

		ShipPosition := model.Pointer{
			PositionX: pointerX,
			PositionY: pointerY,
		}
		ship := model.Ship{Status: 1, Position: ShipPosition}
		player1.Ships = append(player1.Ships, ship)
	}

	coordinatesShip2 := strings.FieldsFunc(inputs[3], Split)
	for _, v := range coordinatesShip2 {
		pointers := strings.Split(v, ",")
		pointerX, err := strconv.Atoi(pointers[0])
		if err != nil {
			log.Fatal(err)
		}

		pointerY, err := strconv.Atoi(pointers[1])
		if err != nil {
			log.Fatal(err)
		}

		ShipPosition := model.Pointer{
			PositionX: pointerX,
			PositionY: pointerY,
		}
		ship := model.Ship{Status: 1, Position: ShipPosition}
		player2.Ships = append(player1.Ships, ship)
	}

	coordinatesMissile1 := strings.FieldsFunc(inputs[5], Split)
	for _, v := range coordinatesMissile1 {
		pointers := strings.Split(v, ",")
		pointerX, err := strconv.Atoi(pointers[0])
		if err != nil {
			log.Fatal(err)
		}

		pointerY, err := strconv.Atoi(pointers[1])
		if err != nil {
			log.Fatal(err)
		}

		MissilePosition := model.Pointer{
			PositionX: pointerX,
			PositionY: pointerY,
		}
		missile := model.Missile{Position: MissilePosition}
		player1.Missile = append(player1.Missile, missile)
	}

	coordinatesMissile2 := strings.FieldsFunc(inputs[6], Split)
	for _, v := range coordinatesMissile2 {
		pointers := strings.Split(v, ",")
		pointerX, err := strconv.Atoi(pointers[0])
		if err != nil {
			log.Fatal(err)
		}

		pointerY, err := strconv.Atoi(pointers[1])
		if err != nil {
			log.Fatal(err)
		}

		MissilePosition := model.Pointer{
			PositionX: pointerX,
			PositionY: pointerY,
		}
		missile := model.Missile{Position: MissilePosition}
		player2.Missile = append(player2.Missile, missile)
	}

	// check attack
	score1 := len(player1.Ships)
	score2 := len(player2.Ships)
	for _, missile1 := range player1.Missile {
		for _, ship2 := range player2.Ships {
			if ship2.Position.PositionX == missile1.Position.PositionX && ship2.Position.PositionY == missile1.Position.PositionY {
				ship2.Status = 0
				score2--
			}
		}
	}

	for _, missile2 := range player2.Missile {
		for _, ship1 := range player1.Ships {
			if ship1.Position.PositionX == missile2.Position.PositionX && ship1.Position.PositionY == missile2.Position.PositionY {
				ship1.Status = 0
				score1--
			}
		}
	}

	fmt.Println(score1)
	fmt.Println(score2)
}
