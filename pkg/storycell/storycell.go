// Package storycell is a collection of structs and functions used in the creation of a TBA [Text Based Adventure]
package storycell

// World is the "One struct to hold them all"
// All the information for the TBA game world
type World struct {
	Cells []MapCell
	Mobs  []Mob
	NPCs  []NPC
	Items []Item
}

// MapCell is the information for each map cell, it is recommended that when
// trying to add special logic, to add a new MapCell and MapConn and then use
// the MapCell's conns to move through the logic.
type MapCell struct {
	ID    float32
	Title string
	Disc  string
	Conns []MapConn
	Mobs  []Mob
	Items []Item
}

// MapConn is a connection between MapCells, it can be called what ever but
// is typically identified by a direction (north, east, south, west) or an
// action/command (talk, grab, pickup, run, attack, etc). You can restrict
// access to a MapConn by adding an item to the recItems slice/array.
type MapConn struct {
	Name      string
	Disc      string
	StartCell MapCell
	EndCell   MapCell
	RecItems  []Item
}

// Mob is an attackable npc, so it will need hp, attack, and defense, if a mob
// is present in a cell it should also have an 'attack' MapCell and MapConn
type Mob struct {
	Name    string
	Disc    string
	HP      int
	Attack  int
	Defense int
	Items   []Item
}

// Item is a item for the Player, NPC, or Mob to use, it adds the hp, attack, and defence
// to the Player, NPC, or Mob respectively
type Item struct {
	Name    string
	Disc    string
	HP      int
	Attack  int
	Defense int
}

// NPC is a non-attackable npc, that may have a dialog or items, for multiple dialogs it
// is recommended for a new version of NPC per dialog.
type NPC struct {
	Name   string
	Dialog string
	Items  []Item
}

// Player is a ... well the player :)
// it also stores the players current values to create a 'save'.
type Player struct {
	Name    string
	Loc     MapCell
	HP      int
	Attack  int
	Defense int
	Inv     []Item
}
