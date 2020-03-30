// Package tbastory is a collection of structs and functions used in the creation of a TBA [Text Based Adventure]
package tbastory

import (
	"encoding/json"
	"fmt"

	"github.com/Gamemastertwig/Gygaria_go/pkg/filewriter"
)

// World is the "One struct to hold them all"
// All the information for the TBA game world
type World struct {
	Cells []MapCell `json:"cell"`
}

// MapCell is the information for each map cell, it is recommended that when
// trying to add special logic, to add a new MapCell and MapConn and then use
// the MapCell's conns to move through the logic.
type MapCell struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Disc  string    `json:"disc"`
	IsEnd bool      `json:"isend"`
	Conns []MapConn `json:"conns"`
	Mobs  []Mob     `json:"mobs"`
	NPCs  []NPC     `json:"npcs"`
	Items []Item    `json:"items"`
}

// MapConn is a connection between MapCells, it can be called what ever but
// is typically identified by a direction (north, east, south, west) or an
// action/command (talk, grab, pickup, run, attack, etc). You can restrict
// access to a MapConn by adding an item to the recItems slice/array.
type MapConn struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	StartCellID   float32 `json:"startcellid"`
	EndCellID     float32 `json:"endcellid"`
	NeedItem      bool    `json:"needitem"`
	NeedNotItem   bool    `json:"neednotitem"`
	RecItemIDs    []int   `json:"recitemids"`
	NeedMobKilled bool    `json:"needmobkilled"`
	NeedMobAlive  bool    `json:"needmobalive"`
	RecMobIDs     []int   `json:"recmobids"`
	NeedNPCTalk   bool    `json:"neednpctalk"`
	NeedNPCNoTalk bool    `json:"neednpcnotalk"`
	RecNPCIDs     []int   `json:"recnpcids"`
}

// Mob is an attackable npc, so it will need hp, attack, and defense, if a mob
// is present in a cell it should also have an 'attack' MapCell and MapConn
type Mob struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Disc    string `json:"disc"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	Items   []Item `json:"items"`
	IsDead  bool   `json:"isdead"`
}

// Item is a item for the Player, NPC, or Mob to use, it adds the hp, attack, and defence
// to the Player, NPC, or Mob respectively
type Item struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Disc    string `json:"disc"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
}

// NPC is a non-attackable npc, that may have a dialog or items, for multiple dialogs it
// is recommended for a new version of NPC per dialog.
type NPC struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Dialog    string `json:"dialog"`
	Items     []Item `json:"items"`
	HasTalked bool   `json:"hastalked"`
}

// Player is a ... well the player :)
// it also stores the players current values to create a 'save'.
type Player struct {
	Name    string `json:"name"`
	PLocID  int    `json:"plocid"`
	CLocID  int    `json:"clocid"`
	HP      int    `json:"hp"`
	Attack  int    `json:"attack"`
	Defense int    `json:"defense"`
	InvIDs  []int  `json:"invids"`
}

// SaveToJSON converts the data in the world and player structs into json and saves them to a file
func SaveToJSON(world World, player Player) {
	worldJSON, err := json.Marshal(world)
	if err != nil {
		fmt.Println(err)
	}
	playerJSON, err := json.Marshal(player)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(string(worldJSON))
	// fmt.Println(string(playerJSON))

	// filewriter.WriteRaw("player.json", playerJSON)
	// filewriter.WriteRaw("world.json", worldJSON)

	filewriter.WriteNew("player.json", playerJSON)
	filewriter.WriteNew("world.json", worldJSON)
}
