package  t3

import("testing")

func TestGameCreation(t *testing.T) {
  players := []Player{Player{Name: "One"}, Player{Name: "Two"}}
  game := Game{Players: players}

  result := game.PlayerCount()

  assert(t, 2, result)
}

func TestPlayerName(t *testing.T) {
  expected1 := "Player One"
  expected2 := "Player Two"
  players := []Player{Player{Name: expected1}, Player{Name: expected2}}
  game := Game{Players: players}

  assert(t, expected1, game.PlayerName(0))
  assert(t, expected2, game.PlayerName(1))
}

func TestValidPlayerMove(t *testing.T) {
  player := Player{Name: "Simon"}
  game := Game{Players: []Player{player}}

  err := game.PlayerMove(player, X, 0)

  assert(t, nil, err)
}

func TestSequentialPlayerMoves(t *testing.T) {
  player1 := Player{Name: "Simon"}
  player2 := Player{Name: "Greg"}
  game := Game{Players: []Player{player1, player2}}

  game.PlayerMove(player1, X, 0)
  err := game.PlayerMove(player1, X, 1)

  assert(t, "Invalid move", err.Error())
}

func TestPlayerMovesInOccupiedCell(t *testing.T) {
  player1 := Player{Name: "Simon"}
  player2 := Player{Name: "Greg"}
  game := Game{Players: []Player{player1, player2}}

  game.PlayerMove(player1, X, 0)
  game.PlayerMove(player2, O, 1)

  err := game.PlayerMove(player1, O, 1)

  assert(t, "Invalid move", err.Error())
}

func assert(t *testing.T, expected interface{}, actual interface{}) {
  if actual != expected {
    t.Fatalf("Expected: %v\n\tGot: %v", expected, actual)
  }
}
