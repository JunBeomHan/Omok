package board

import "fmt"

const (
	BOARD_SIZE = 15
)

const (
	EMPTY      = 0
	ONE_PLAYER = 1
	TWO_PLAYER = 0
)

type Board struct {
	board [BOARD_SIZE][BOARD_SIZE]int
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) initBoard() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			b.board[i][j] = 0
		}
	}
}

// CheckStonePresence는 바둑돌이 있으면 1, 없으면 0을 반환합니다.
func (b *Board) checkStonePresence(row, colum int) bool {
	return b.board[row][colum] != EMPTY
}

func (b *Board) printBoard() {
	fmt.Print("  ")
	for i := 0; i < BOARD_SIZE; i++ {
		if i < 10 {
			fmt.Printf("%2d ", i)
		} else {
			fmt.Printf("%3d", i)
		}
	}
	fmt.Println()
	for i := 0; i < BOARD_SIZE; i++ {
		fmt.Printf("%2d ", i)
		for j := 0; j < BOARD_SIZE; j++ {
			if b.board[i][j] == EMPTY {
				fmt.Print("○  ")
			} else if b.board[i][j] == ONE_PLAYER {
				fmt.Print("①  ")
			} else if b.board[i][j] == TWO_PLAYER {
				fmt.Print("②  ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) checkWinner(s Stone, row, colum int) bool {

	// var case int = 0
	// 4 가지

	// case 1 ㅡ
	for i := colum - 2; i <= colum+2; i++ {
		if (colum-2 < 0) || (colum+2 > 14) {
			break
		}
		if b.board[row][i] != s.stoneStatus {
			break
		}
		return true
	}

	// case 2 |
	for i := row - 2; i <= row+2; i++ {
		if (row-2 < 0) || (row+2 > 14) {
			break
		}
	}
}
