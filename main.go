package main

import (
    "fmt"
    "os"

    "./models"
)

func main() {
    game := models.CreateGame(3)
    player1, err := models.NewPlayer("player#1", 1)
    if  err != nil {
        fmt.Errorf("error in creating player")
        os.Exit(1)
    }
    player2, err := models.NewPlayer("player#1", 2)
    if  err != nil {
        fmt.Errorf("error in creating player")
        os.Exit(1)
    }
    _ = game.AddPlayer(*player1)
    _ = game.AddPlayer(*player2)
    game.Play(game.GetNextPlayer(), 0, 0)
    game.Play(game.GetNextPlayer(), 1, 0)
    game.Play(game.GetNextPlayer(), 1, 1)
    game.Play(game.GetNextPlayer(), 1, 1)
    game.Play(game.GetNextPlayer(), 2, 2)

    game.Display()
    println(game.IsGameFinished())
}
