package connect

import "strings"

const (
	dot = iota
	x
	xVisited
	o
	oVisited
)

type Board [][]int
type Pos [2]int

func ResultOf(lines []string) (string, error) {
	board := make(Board, 0)
	for _, s := range lines {
		boardLine := make([]int, 0)
		s = strings.ReplaceAll(s, " ", "")
		for _, r := range s {
			switch r {
			case '.':
				boardLine = append(boardLine, dot)
			case 'X':
				boardLine = append(boardLine, x)
			case 'O':
				boardLine = append(boardLine, o)
			}
		}
		board = append(board, boardLine)
	}

	if len(board) > 0 {
		//test O's starting from top
		for x, _ := range board[0] {
			if board.solve(o, Pos{0, x}) {
				return "O", nil
			}
		}

		//test X's starting from left
		for y, _ := range board {
			if board.solve(x, Pos{y, 0}) {
				return "X", nil
			}
		}
	}
	return "", nil
}

func (b Board) solve(mark int, pos Pos) (rez bool) {
	if b.inBoard(pos) {
		if b[pos[0]][pos[1]] == dot {
			return false
		}
		if b[pos[0]][pos[1]] == mark {
			//mark field as visited
			b[pos[0]][pos[1]] = mark + 1
			//check if bottom is reached in case of 'O'
			if mark == o && pos[0] == len(b)-1 {
				return true
			}
			//check if right hand side is reached in case of 'X'
			if mark == x && pos[1] == len(b[0])-1 {
				return true
			}
			if b.solve(mark, Pos{pos[0], pos[1] + 1}) ||
				b.solve(mark, Pos{pos[0], pos[1] - 1}) ||
				b.solve(mark, Pos{pos[0] + 1, pos[1]}) ||
				b.solve(mark, Pos{pos[0] + 1, pos[1] - 1}) ||
				b.solve(mark, Pos{pos[0] - 1, pos[1]}) ||
				b.solve(mark, Pos{pos[0] - 1, pos[1] + 1}) {
				return true
			}
		}
	}
	return false
}

func (b Board) inBoard(pos Pos) bool {
	if pos[0] < 0 || pos[1] < 0 || pos[0] >= len(b) || pos[1] >= len(b[0]) {
		return false
	}
	return true
}

func (b Board) bottomOfBoard(pos Pos) bool {
	if pos[1] == len(b)-1 {
		return true
	}
	return false
}
