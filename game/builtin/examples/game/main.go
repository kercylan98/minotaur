package main

import "minotaur/game/builtin/examples/game/game"

func main() {
	if err := game.Server.Run(":9999"); err != nil {
		panic(err)
	}
}
