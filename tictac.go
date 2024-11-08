package main

import (
	"fmt"
	"strconv"
	"strings"
)

func turn(board [][]string, player string) {
    fmt.Println("Player", player)
    for i := 0; i < len(board); i++ {
        fmt.Println(board[i])
    }
    fmt.Println("Enter row and column (0, 1, or 2):")

    var row, col int
    for {
        var input string
        fmt.Scanln(&input)
        parts := strings.Split(input, " ")
        if len(parts) != 2 {
            fmt.Println("Invalid input. Please enter row and column separated by a space.")
            continue
        }

        var err error
        row, err = strconv.Atoi(parts[0])
        if err != nil || row < 0 || row > 2 {
            fmt.Println("Invalid row. Please enter a number between 0 and 2.")
            continue
        }

        col, err = strconv.Atoi(parts[1])
        if err != nil || col < 0 || col > 2 {
            fmt.Println("Invalid column. Please enter a number between 0 and 2.")
            continue
        }

        if board[row][col] != "_" {
            fmt.Println("Cell already occupied, try again.")
            continue
        }

        break
    }

    board[row][col] = player
}

func checkWin(board [][]string, player string) bool {
    // Check rows and columns
    for i := 0; i < 3; i++ {
        if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||
            (board[0][i] == player && board[1][i] == player && board[2][i] == player) {
            return true
        }
    }
    // Check diagonals
    if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||
        (board[0][2] == player && board[1][1] == player && board[2][0] == player) {
        return true
    }
    return false
}

func checkDraw(board [][]string) bool {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if board[i][j] == "_" {
                return false
            }
        }
    }
    return true
}

type Player struct {
    name string
}

func main() {
    board := [][]string{
        {"_", "_", "_"},
        {"_", "_", "_"},
        {"_", "_", "_"},
    }

    player1 := Player{"X"}
    player2 := Player{"O"}

    gameOver := false
    currentPlayer := player1

    for !gameOver {
        turn(board, currentPlayer.name)
        if checkWin(board, currentPlayer.name) {
            fmt.Println("Player", currentPlayer.name, "wins!")
            gameOver = true
        } else if checkDraw(board) {
            fmt.Println("It's a draw!")
            gameOver = true
        } else {
            if currentPlayer.name == player1.name {
                currentPlayer = player2
            } else {
                currentPlayer = player1
            }
        }
    }
}