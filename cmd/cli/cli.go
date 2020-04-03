package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/Gamemastertwig/Gygaria_go/pkg/userinputs"

	"github.com/Gamemastertwig/Gygaria_go/cmd/gygaria"
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
)

var clear map[string]func() // create a map for storing clear functions

// game/engine variables
var world tbastory.World
var player tbastory.Player

func init() {
	world = gygaria.CreateGygaria()
	player = gygaria.CreatePlayer()

	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	displayCurrentCell()
}

func callClear() {
	function, good := clear[runtime.GOOS]
	if good {
		function()
	} else {
		fmt.Println("Platform not supported, cannot clear terminal screen!")
	}
}

func displayCurrentCell() {
	callClear()
	currentCellID := player.CLocID

	for _, cell := range world.Cells {
		if cell.ID == currentCellID {
			fmt.Println(cell.Title)
			fmt.Println(cell.Disc)
			fmt.Println("")
			var options []string
			var optionIds []int
			for _, conn := range cell.Conns {
				// control variables
				connGood := false
				pass := 0
				count := 0

				// test variables
				temp := ""

				if conn.NeedItem {
					// check inv for items in conn
					for _, itemID := range conn.RecItemIDs {
						count++ // check each required item once
						if checkHaveItem(itemID) {
							pass++ // if item in inv

							// test logic [remove later]
							temp = temp + "(" + getItemName(itemID) + " is in inventory)"
						}
					}
				}
				if conn.NeedNotItem {
					// check inv for the absense of items in conn
					for _, itemID := range conn.RecItemIDs {
						count++ // check each required item once
						if !checkHaveItem(itemID) {
							pass++ // if item not in inv

							// test logic [remove later]
							temp = temp + "(" + getItemName(itemID) + " is not in inventory)"
						}
					}
				}
				if conn.NeedMobAlive {
					// check that mob is alive for mobs in conn
					for _, recMobID := range conn.RecMobIDs {
						count++ // check each required mob once
						if !checkMobDead(recMobID) {
							pass++ // if mob is alive

							// test logic [remove later]
							temp = temp + "(" + getMobName(recMobID) + " is alive)"
						}
					}
				}
				if conn.NeedMobKilled {
					// check that mob is dead for mobs in conn
					for _, recMobID := range conn.RecMobIDs {
						count++ // check each required mob once
						if checkMobDead(recMobID) {
							pass++ // if mob is dead

							// test logic [remove later]
							temp = temp + "(" + getMobName(recMobID) + " is dead)"
						}
					}
				}
				if conn.NeedNPCTalk {
					// check that npc talked for npcs in conn+
					for _, recNpcID := range conn.RecNPCIDs {
						count++ // check each required NPC once
						if checkNpcTalked(recNpcID) {
							pass++ // if npc has talked

							// test logic [remove later]
							temp = temp + "(" + getNpcName(recNpcID) + " has talked)"
						}
					}
				}
				if conn.NeedNPCNoTalk {
					// check that npc have not talked for npcs in conn
					for _, recNpcID := range conn.RecNPCIDs {
						count++ // check each required NPC once
						if !checkNpcTalked(recNpcID) {
							pass++ // if npc has not talked

							// test logic [remove later]
							temp = temp + "(" + getNpcName(recNpcID) + " has not talked)"
						}
					}
				}

				// TESTING
				countStr := "Count: " + strconv.Itoa(count)
				passStr := "Pass: " + strconv.Itoa(pass)
				fmt.Println(countStr + " " + passStr + " " + temp + " [" + conn.Name + "]")

				// if pass and count are equal then they are a valid/confirmed connection
				if pass == count {
					connGood = true
				}

				// if the connection is good add it to the players options
				if connGood {
					options = append(options, conn.Name)
					optionIds = append(optionIds, conn.EndCellID)
				}
			}
			choice := userinputs.MultiChoiceAnswer("What would you like to do?", options)
			if choice == "TALK" {
				// do talk things here
			} else if choice == "ATTACK" {
				// do attack things here
			} else {
				for i, op := range options {
					if op == choice {
						player.PLocID = player.CLocID
						player.CLocID = optionIds[i]
						fmt.Printf("Player.CLocID: %v ", player.CLocID)
						displayCurrentCell()
					}
				}
			}
			fmt.Println(choice)
		}
	}
	return
}

func checkHaveItem(itemID int) bool {
	// check if item is in inv
	for _, invID := range player.InvIDs {
		// if item present in inv return true
		if invID == itemID {
			return true
		}
	}
	// otherwise return false
	return false
}

func getItemName(itemID int) string {
	// check each cell
	for _, cell := range world.Cells {
		// check items in each cell
		for _, cellItem := range cell.Items {
			// compaire item(s) with one we are looking for
			if cellItem.ID == itemID {
				// if match return cellItem.Name
				return cellItem.Name
			}
		}

		// check each mob in each cell (if not found in cell directly)
		for _, mob := range cell.Mobs {
			// check each item on the mobs
			for _, mobItem := range mob.Items {
				// compaire item(s) with one we are looking for
				if mobItem.ID == itemID {
					// if match return cellItem.Name
					return mobItem.Name
				}
			}
		}

		// check each npc in each cell (if not found in cell or mobs directly)
		for _, npc := range cell.NPCs {
			// check each item on the mobs
			for _, npcItem := range npc.Items {
				// compaire item(s) with one we are looking for
				if npcItem.ID == itemID {
					// if match return cellItem.Name
					return npcItem.Name
				}
			}
		}
	}
	// otherwise return blank
	return ""
}

func checkMobDead(mobID int) bool {
	// check each cell
	for _, cell := range world.Cells {
		// check mob(s) in each cell
		for _, mob := range cell.Mobs {
			// comaire mob(s) with the one we are checking
			if mob.ID == mobID {
				// if they are dead return true
				if mob.IsDead {
					return true
				}
			}
		}
	}
	// otherwise return false
	return false
}

func getMobName(mobID int) string {
	// check each cell
	for _, cell := range world.Cells {
		// check mob(s) in each cell
		for _, mob := range cell.Mobs {
			// comaire mob(s) with the one we are checking
			if mob.ID == mobID {
				// if they are the same return mob.Name
				return mob.Name
			}
		}
	}
	// otherwise return blank
	return ""
}

func checkNpcTalked(npcID int) bool {
	// check each cell
	for _, cell := range world.Cells {
		// check npc(s) in each cell
		for _, npc := range cell.NPCs {
			// comaire npc(s) with the one we are checking
			if npc.ID == npcID {
				// if they talked return true
				if npc.HasTalked {
					return true
				}
			}
		}
	}
	// otherwise return false
	return false
}

func getNpcName(npcID int) string {
	// check each cell
	for _, cell := range world.Cells {
		// check npc in each cell
		for _, npc := range cell.NPCs {
			// comaire npcs with the one where are checking
			if npc.ID == npcID {
				// if they match return npc.Name
				return npc.Name
			}
		}
	}
	// otherwise return blank
	return ""
}
