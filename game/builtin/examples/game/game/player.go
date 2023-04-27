package game

import (
	"log"
	"minotaur/game"
	"minotaur/game/builtin"
	"minotaur/server"
	"time"
)

func NewPlayer(id int64, conn *server.Conn) *Player {
	player := &Player{
		Player: builtin.NewPlayer[int64](id, conn),
	}
	gameplay := builtin.NewGameplay()
	gameplayOver := builtin.NewGameplayOver()
	gameplayTime := builtin.NewGameplayTime(gameplay, gameplayOver)

	player.GameplayTime = gameplayTime

	return player
}

type Player struct {
	*builtin.Player[int64]
	game.GameplayTime
}

func (slf *Player) Start() {
	_ = slf.GameStart(func() error {
		log.Println("game start, init map...")
		return nil
	})
}

func (slf *Player) onGameplayStart(t time.Time) {
	slf.SetEndTime(t.Add(60 * time.Second))
	log.Println("the game will end in 60 seconds")
}

func (slf *Player) onGameplayOver() {
	log.Println("game over")
}
