package main

import (
	"github.com/Gamemastertwig/Gygaria_go/cmd/gygaria"
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
)

func main() {
	world := gygaria.CreateGygaria()
	player := gygaria.CreatePlayer()

	// fmt.Println(world)
	// fmt.Println(player)

	tbastory.SaveToJSON(world, player)
}
