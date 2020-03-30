package main

import (
	"github.com/Gamemastertwig/Gygaria_go/cmd/gygaria"
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
)

func main() {
	world := gygaria.CreateGygaria()
	player := gygaria.CreatePlayer()

	tbastory.WorldToJSON(world)
	tbastory.PlayerToJSON(player)

	// world := tbastory.JSONToWorld("world.json")
	// player := tbastory.JSONToPlayer("player.json")

	// fmt.Println(world)
	// fmt.Println(player)
}
