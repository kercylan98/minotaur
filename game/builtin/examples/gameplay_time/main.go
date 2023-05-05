package main

import (
	"log"
	"minotaur/game"
	"minotaur/game/builtin"
	"sync"
	"time"
)

func NewGame() *Game {
	gameplay := builtin.NewGameplay()
	gameplayOver := builtin.NewGameplayOver()
	return &Game{
		GameplayTime: builtin.NewGameplayTime(gameplay, gameplayOver),
	}
}

type Game struct {
	game.GameplayTime
	wait sync.WaitGroup
}

func (slf *Game) init() error {
	slf.wait.Add(1)
	return nil
}

func (slf *Game) onGameStart(startTime time.Time) {
	log.Println("游戏开始")
	slf.SetEndTime(startTime.Add(3 * time.Second))
}

func (slf *Game) onGameOver() {
	log.Println("游戏结束")
	slf.wait.Done()
}

func main() {
	g := NewGame()
	g.RegGameplayStartEvent(g.onGameStart)
	g.RegGameplayOverEvent(g.onGameOver)

	if err := g.GameStart(g.init); err != nil {
		panic(err)
	}

	g.wait.Wait()
}
