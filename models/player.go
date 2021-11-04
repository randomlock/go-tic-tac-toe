package models

import "fmt"

type Player struct {
    number int
    name string
}

func NewPlayer(playerName string, playerNumber int) (*Player, error)  {
    if playerNumber < 1 || playerNumber > 2 {
        return nil, fmt.Errorf("incorrect player number")
    }
    return &Player{
        number: playerNumber,
        name:   playerName,
    }, nil
}