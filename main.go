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

	// initialize battlefield
	size, err := strconv.Atoi(inputs[0])
	if err != nil {
		log.Fatal(err)
	}

	// initialize matrix
	fieldPlayer1 := make([][]string, size)
	for i := 0; i < size; i++ {
		fieldPlayer1[i] = make([]string, size)
	}

	fieldPlayer2 := make([][]string, size)
	for i := 0; i < size; i++ {
		fieldPlayer2[i] = make([]string, size)
	}

	// initialize player
	player1 := model.Player{PlayerID: 1}
	player2 := model.Player{PlayerID: 2}

	// add coordinates of ships
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
		ship := model.Ship{Status: 2, Position: ShipPosition}
		player1.Ships = append(player1.Ships, ship)
		fieldPlayer1[pointerX][pointerY] = "B"
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
		ship := model.Ship{Status: 2, Position: ShipPosition}
		player2.Ships = append(player2.Ships, ship)
		fieldPlayer2[pointerX][pointerY] = "B"
	}

	// add coordinates of misilles
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
		fieldPlayer2[pointerX][pointerY] = "O"
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
		fieldPlayer1[pointerX][pointerY] = "O"
	}

	// check attack usecase
	score1 := len(player1.Ships)
	score2 := len(player2.Ships)
	for _, missile1 := range player1.Missile {
		for index, ship2 := range player2.Ships {
			if ship2.Position.PositionX == missile1.Position.PositionX && ship2.Position.PositionY == missile1.Position.PositionY {
				player2.Ships[index].Status--
				if player2.Ships[index].Status == 0 {
					fieldPlayer2[ship2.Position.PositionX][ship2.Position.PositionY] = "X"
					score2--
				}
			}
		}
	}

	for _, missile2 := range player2.Missile {
		for index, ship1 := range player1.Ships {
			if ship1.Position.PositionX == missile2.Position.PositionX && ship1.Position.PositionY == missile2.Position.PositionY {
				player1.Ships[index].Status--
				if player1.Ships[index].Status == 0 {
					fieldPlayer1[ship1.Position.PositionX][ship1.Position.PositionY] = "X"
					score1--
				}
			}
		}
	}

	fmt.Println("Score Player 1:", score1)
	fmt.Println("Score Player 2:", score2)
	fmt.Println("Board Player 1:")
	for _, v := range fieldPlayer1 {
		fmt.Println(v)
	}

	fmt.Println("Board Player 2:")
	for _, v := range fieldPlayer2 {
		fmt.Println(v)
	}

	if score1 == score2 {
		fmt.Println("Result: Draw")
	} else if score1 > score2 {
		fmt.Println("Result: Player 1 Win")
	} else {
		fmt.Println("Result: Player 2 Win")
	}
}
