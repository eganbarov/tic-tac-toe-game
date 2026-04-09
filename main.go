package main

import (
	"errors"
	"fmt"
)

type Board [3][3]int
type Player map[string]int
type MoveState int

func main() {
	board := Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	playerMap := Player{
		"X": 1,
		"O": 2,
	}
	var moveState MoveState

	var row, column int
	var player string
	var isGameFineshed bool = false

	fmt.Println("Player 1: X; Player 2: O")

	for !isGameFineshed {
		drawBoard(&board)

		fmt.Println("Input row, column and value:")
		_, err := fmt.Scanln(&row, &column, &player)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		err = checkInputParams(row, column, player)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		err = checkCurrentMove(playerMap, player, moveState)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		err = checkCell(&board, row, column)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		currentPlayer := playerMap[player]
		board[row][column] = currentPlayer
		moveState = MoveState(currentPlayer)

		if checkWin(&board) {
			isGameFineshed = true
			drawBoard(&board)
		}
	}

	fmt.Println("Congrats, the game is finished, launch a new one!")
}

func checkInputParams(row, column int, player string) error {
	if row < 0 || row > 2 {
		return errors.New("Row can be from 0 to 2, choose value from correct range")
	}

	if column < 0 || column > 2 {
		return errors.New("Column can be from 0 to 2, choose value from correct range")
	}

	if player != "X" && player != "O" {
		return errors.New("Value can be from X or O, choose the correct one")
	}

	return nil
}

func drawBoard(b *Board) {
	fmt.Println("")
	for i := range 3 {
		for j := range 3 {
			switch b[i][j] {
			case 1:
				fmt.Print("X")
			case 2:
				fmt.Print("O")
			default:
				fmt.Print("-")
			}
			if j < 2 {
				fmt.Print(" | ")
			}
		}
		fmt.Println("")
		if i < 2 {
			fmt.Println("----------")
		}
	}
	fmt.Println("")
}

func checkCurrentMove(playerMap Player, player string, moveState MoveState) error {
	currentPlayer := playerMap[player]
	if moveState == MoveState(currentPlayer) {
		return errors.New("Now it's not your move")
	}

	return nil
}

func checkCell(b *Board, row, column int) error {
	if (*b)[row][column] != 0 {
		return errors.New("This cell has already used")
	}

	return nil
}

func checkWin(b *Board) bool {
	for i := range 3 {
		if b[i][0] != 0 && b[i][0] == b[i][1] && b[i][0] == b[i][2] {
			return true
		}
		if b[0][i] != 0 && b[0][i] == b[1][i] && b[0][i] == b[2][i] {
			return true
		}
	}

	if b[0][0] != 0 && b[0][0] == b[1][1] && b[0][0] == b[2][2] {
		return true
	}

	if b[0][2] != 0 && b[0][2] == b[1][1] && b[0][2] == b[2][0] {
		return true
	}

	return false
}
