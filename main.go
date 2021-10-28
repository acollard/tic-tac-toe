package main

import (
	"fmt"
)

func main() {

	tracker := make(map[string]string)

	/* Gameboard
	   0 1 2
	   3 4 5
	   6 7 8
	*/
	gameboard := make([]string, 9)
	play(tracker, gameboard, 0)
	xWins := 0
	yWins := 0
	draws := 0
	for _, v := range tracker {
		if v == "X" {
			xWins++
		} else if v == "O" {
			yWins++
		} else if v == "D" {
			draws++
		}
	}
	fmt.Printf("X wins: %v, Y wins: %v, Draws: %v", xWins, yWins, draws)
}

func play(tracker map[string]string, gameboard []string, turn int) {
	if turn < 9 { // is the game over
		turn++
		for i := 0; i < 9; i++ {
			if gameboard[i] == "" {
				boardCopy := make([]string, 9)
				copy(boardCopy, gameboard)
				if turn%2 == 1 {
					boardCopy[i] = "X"
				} else {
					boardCopy[i] = "O"
				}

				hash := getHash(boardCopy)
				if tracker[hash] != "" { // we've seen this game before so we can exit
					continue
				}
				winner := ""
				if turn >= 5 {
					// Check for a winner
					winner = whoWon(boardCopy)
				}
				if winner != "" {
					tracker[hash] = winner
				} else {
					// no winner yet
					play(tracker, boardCopy, turn)
				}
			}
		}
	} else {
		hash := getHash(gameboard)
		tracker[hash] = "D"
	}
}

func getHash(board []string) string {
	hash := ""
	for _, v := range board {
		if v == "" {
			hash += " "
		} else {
			hash += v
		}
	}
	return hash
}

// there are 8 unique win locations for each side
func whoWon(board []string) string {
	current := board[0]
	if board[1] == current && board[2] == current {
		return current
	}
	if board[3] == current && board[6] == current {
		return current
	}
	if board[4] == current && board[8] == current {
		return current
	}
	current = board[1]
	if board[4] == current && board[7] == current {
		return current
	}
	current = board[2]
	if board[4] == current && board[6] == current {
		return current
	}
	if board[5] == current && board[8] == current {
		return current
	}
	current = board[3]
	if board[4] == current && board[5] == current {
		return current
	}
	current = board[6]
	if board[7] == current && board[8] == current {
		return current
	}
	return ""
}
