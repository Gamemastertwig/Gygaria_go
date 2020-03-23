package storycell

// World is the "One struct to hold them all"
// All the information for the TBA game world
type World struct {
	cells []MapCell
	mobs  []Mob
	npcs  []NPC
	items []Item
}

// MapCell is the information for each map cell, it is recommended that when
// trying to add special logic, to add a new MapCell and MapConn and then use
// the MapCell's conns to move through the logic, also you can restrict access
// to a MapCell by adding an item to the recItems slice/array
type MapCell struct {
	id       float32
	title    string
	disc     string
	conns    []MapConn
	mobs     []Mob
	items    []Item
	recItems []Item
}

// MapConn is a connection between MapCells, it can be called what ever but is
// typically identified by a direction (north, east, south, west) or an
// action/command (talk, grab, pickup, run, attack, etc).
type MapConn struct {
	name      string
	disc      string
	startCell MapCell
	endCell   MapCell
}

// Mob is an attackable npc, so it will need hp, attack, and defense, if a mob
// is present in a cell it should also have an 'attack' MapCell and MapConn
type Mob struct {
	name    string
	disc    string
	hp      int
	attack  int
	defense int
}

// Item is a item for the Player, NPC, or Mob to use, it adds the hp, attack, and defence
// to the Player, NPC, or Mob respectively
type Item struct {
	name string
	disc string
}

// NPC is a non-attackable npc, that may have a dialog or items, for multiple dialogs it
// is recommended for a new version of NPC per dialog.
type NPC struct {
	name   string
	dialog string
	items  []Item
}

// Player is a ...
type Player struct {
	name    string
	loc     MapCell
	hp      int
	attack  int
	defense int
	inv     []Item
}
