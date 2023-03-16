package stone

type Stone struct {
	stoneStatus int
	isWin       bool
}

func NewStone() Stone {
	return Stone{}
}

func (s *Stone) GetStoneStatus() int {
	return s.stoneStatus
}

func (s *Stone) GetIsWin() bool {
	return s.isWin
}

func (s *Stone) placeMove(b *Board, row, colum int) {
	switch s.stoneStatus {
	case ONE_PLAYER:
		b.board[row][colum] = ONE_PLAYER
	case TWO_PLAYER:
		b.board[row][colum] = TWO_PLAYER
	}

}
