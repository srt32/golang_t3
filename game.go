package  t3

import(
  "errors"
)

const EMPTY_CELL_VALUE int = 0
const O int = 1
const X int = 2

type Player struct {
  Name string
}

type Game struct {
  Players []Player
  playCount int
  cells [9]int
}

func (g *Game) PlayerCount() int {
  return len(g.Players)
}

func (g *Game) PlayerName(index int) string {
  return g.Players[index].Name
}

func (g *Game) PlayerMove(player Player, token int, cellIndex int) (result error) {
  if g.cells[cellIndex] == EMPTY_CELL_VALUE && g.canPlayerPlay(player) {
    g.cells[cellIndex] = token
    g.playCount++
  } else {
    result = errors.New("Invalid move")
  }

  return result
}

func (g *Game) canPlayerPlay(player Player) bool {
  for index, p := range g.Players {
    if p == player {
      return g.playCount % 2 == index
    }
  }

  return false
}
