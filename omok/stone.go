package omok

type Stone struct {
	stoneStatus int
	Row         int
	Colum       int
}

func NewStone(stoneStatus int) Stone {
	return Stone{stoneStatus: stoneStatus}
}

func (s *Stone) GetStoneStatus() int {
	return s.stoneStatus
}

func (s *Stone) SetStoneStatus(stoneStatus int) {
	s.stoneStatus = stoneStatus
}

// placemove 하는 자리에 돌이 존재하면 안됌

func (s *Stone) checkDanger(b *Board) bool {
	// 인덱스 범위 벗어나면 안됌
	if s.Colum > 14 && s.Row > 14 || b.board[s.Colum][s.Row] != 0 {
		return false
	}
	return true
}

func (s *Stone) PlaceMove(b *Board) bool {

	if s.checkDanger(b) {
		switch s.stoneStatus {
		case ONE_PLAYER:
			b.board[s.Row][s.Colum] = ONE_PLAYER
		case TWO_PLAYER:
			b.board[s.Row][s.Colum] = TWO_PLAYER
		}
		return true
	}
	return false

}
