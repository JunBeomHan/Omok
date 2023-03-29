package omok

import "fmt"

const (
	BOARD_SIZE = 15
)

const (
	EMPTY      = 0
	ONE_PLAYER = 1
	TWO_PLAYER = 2
)

type Board struct {
	board [BOARD_SIZE][BOARD_SIZE]int
}

func NewBoard() Board {
	return Board{}
}

func (b *Board) InitBoard() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			b.board[i][j] = 0
		}
	}
}

func (b *Board) PrintBoard() {
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

// CheckStonePresence는 바둑돌이 있으면 1, 없으면 0을 반환합니다.
func (b *Board) CheckStonePresence(s *Stone) bool {
	return b.board[s.Colum][s.Row] != EMPTY
}

func (b *Board) CheckWinner(s *Stone) bool {
	if isCaseOne(b, s) || isCaseTwo(b, s) || isCaseThree(b, s) || isCaseFour(b, s) || isCaseSix(b, s) || isCaseSeven(b, s) {
		return true
	}
	return false
}

func isCaseOne(b *Board, s *Stone) bool {

	if (s.Row - 4) < 0 {
		return false
	}

	for i := s.Row - 4; i <= s.Row; i++ {
		if s.stoneStatus != b.board[s.Colum][i] {
			return false
		}
	}
	return true

}

func isCaseTwo(b *Board, s *Stone) bool {

	if (s.Row + 4) > 14 {
		return false
	}

	for i := s.Row; i <= s.Row+4; i++ {
		if s.stoneStatus != b.board[s.Colum][i] {
			return false
		}
	}
	return true
}

func isCaseThree(b *Board, s *Stone) bool {

	if (s.Colum - 4) < 0 {
		return false
	}

	for i := s.Colum - 4; i <= s.Colum; i++ {
		if s.stoneStatus != b.board[i][s.Row] {
			return false
		}
	}
	return true
}

func isCaseFour(b *Board, s *Stone) bool {

	if (s.Colum + 4) > 14 {
		return false
	}

	for i := s.Colum; i <= s.Colum+4; i++ {
		if s.stoneStatus != b.board[i][s.Row] {
			return false
		}
	}
	return true
}

func isCaseSix(b *Board, s *Stone) bool {
	if s.Colum+4 > 14 || s.Row+4 > 14 {
		return false
	}
	// \ case

	colum, row := s.Colum, s.Row
	for i := 1; i <= 4; i++ {
		colum++
		row++
		if s.stoneStatus != b.board[colum][row] {
			return false
		}
	}
	return true
}

func isCaseSeven(b *Board, s *Stone) bool {
	if s.Colum-4 < 0 || s.Row-4 < 0 {
		return false
	}

	colum, row := s.Colum, s.Row
	for i := 1; i <= 4; i++ {
		colum--
		row--
		if s.stoneStatus != b.board[colum][row] {
			return false
		}
	}
	return true
}
