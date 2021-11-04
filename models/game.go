package models

import "fmt"

var playerSymbol  = map[int]string{
    1: "X",
    2: "0",
}

type Game struct {
    player []Player
    size int
    grid [][]string
    currentPlayer Player
}

func CreateGame(size int) *Game {
    
    game := &Game{
        size:    size,
    }
    game.grid = make([][]string, size)
    for i := range game.grid {
        game.grid[i] = make([]string, size)
    }
    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            game.grid[i][j] = " "
        }
    }
    return game
}

func (game *Game) Display()  {
    // game.grid[0][0] = "X"
    // game.grid[0][1] = "0"
    // game.grid[0][2] = "0"
    // game.grid[1][0] = "X"
    // game.grid[1][1] = "0"
    // game.grid[1][2] = "0"
    // game.grid[2][0] = "X"
    // game.grid[2][1] = "0"
    // game.grid[2][2] = "0"
    for i := 0; i < game.size; i++ {
        for j := 0; j < game.size; j++ {
            fmt.Print(" ")
            fmt.Print( game.grid[i][j] + "|")
            fmt.Print(" ")
        }
        endStr := ""
        for k := 0; k < game.size * ( 1 + game.size); k++ {
            endStr += "-"
        }
        println()
        println(endStr)
    }
}

func (game *Game) GetNextPlayer() Player {
    if game.player[0] == game.currentPlayer {
        game.currentPlayer  = game.player[1]
        return game.player[0]
    } else {
        game.currentPlayer  = game.player[0]
        return game.player[1]
    }
}

func (game *Game) AddPlayer(player Player) (err error)  {
    if len(game.player) >= 2 {
        return fmt.Errorf("cannot add more than 2 player")
    }
    game.player = append(game.player, player)
    return
}

func (game *Game) Play(player Player, x int, y int) (err error)  {
    if x < 0 || x >= game.size {
        game.setCurrentPlayer(player)
        return fmt.Errorf("invalid box")
    } 
    if y <0 || y >= game.size {
        game.setCurrentPlayer(player)
        return fmt.Errorf("invalid box")
    }
    if game.grid[x][y] != " " {
        game.setCurrentPlayer(player)
        return fmt.Errorf("this box is already used")
    }
    game.grid[x][y] = playerSymbol[player.number]
    return
}

func (game *Game) setCurrentPlayer(player Player)  {
    game.currentPlayer = player
}

func (game *Game) IsGameFinished() bool {
    // check horizontally
    for i := 0; i < game.size; i++ {
        symbol := game.grid[i][0]
        hasWon := true
        for j := 0; j < game.size; j++ {
            if game.grid[i][j] != symbol || game.grid[i][j] == " " {
                hasWon = false
                break
            }
        }
        if hasWon {
            return true
        }
    }

    // check vertically
    for i := 0; i < game.size; i++ {
        symbol := game.grid[0][i]
        hasWon := true
        for j := 0; j < game.size; j++ {
            if game.grid[j][i] != symbol || game.grid[j][i] == " " {
                hasWon = false
                break
            }
        }
        if hasWon {
            return true
        }
    }

    // check diagonal #1
    symbol := game.grid[0][0]
    hasWon := true
    for i := 0; i < game.size; i++ {
        if game.grid[i][i] != symbol || game.grid[i][i] == " " {
            hasWon = false
            break
        }
    }
    if hasWon {
        return true
    }

    // check diagonal #2
    symbol = game.grid[0][game.size-1]
    for i := game.size-1; i >= 0; i-- {
        if game.grid[i][i] != symbol || game.grid[i][i] == " " {
            hasWon = false
            break
        }
    }
    if hasWon {
        return true
    }

    return false
}