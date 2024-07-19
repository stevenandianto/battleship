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
}
