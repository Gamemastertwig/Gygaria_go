package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/Gamemastertwig/Gygaria_go/cmd/gygaria"
	"github.com/Gamemastertwig/Gygaria_go/pkg/randimizer"
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
	"github.com/Gamemastertwig/Gygaria_go/pkg/userinputs"
)

var clear map[string]func() // create a map for storing clear functions

// game/engine variables
var world tbastory.World
var player tbastory.Player
var endGame bool

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
	for !endGame {
		displayCurrentCell()
	}
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
			var speakerID int
			var speaker string
			var spkDialog string
			var monsterIDs []int

			// testing
			fmt.Println("| DEBUG Start |")

			for _, conn := range cell.Conns {
				// control variables
				connGood := false
				pass := 0
				count := 0

				// test variables
				// temp := ""

				if conn.NeedItem {
					// check inv for items in conn
					for _, itemID := range conn.RecItemIDs {
						count++ // check each required item once
						if checkHaveItem(itemID) {
							pass++ // if item in inv

							// test logic [remove later]
							// temp = temp + "(" + getItemName(itemID) + " is in inventory)"
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
							// temp = temp + "(" + getItemName(itemID) + " is not in inventory)"
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
							// temp = temp + "(" + getMobName(recMobID) + " is alive)"
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
							// temp = temp + "(" + getMobName(recMobID) + " is dead)"
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
							// temp = temp + "(" + getNpcName(recNpcID) + " has talked)"
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
							// temp = temp + "(" + getNpcName(recNpcID) + " has not talked)"
						}
					}
				}

				// TESTING
				// countStr := "Count: " + strconv.Itoa(count)
				// passStr := "Pass: " + strconv.Itoa(pass)
				// fmt.Println("| " + countStr + " " + passStr + " " + temp + " [" + conn.Name + "] |")

				// if pass and count are equal then they are a valid/confirmed connection
				if pass == count {
					connGood = true
				}

				// if the connection is good add it to the players options
				if connGood {
					options = append(options, conn.Name)
					optionIds = append(optionIds, conn.EndCellID)
				}

				if conn.Name == "TALK" {
					// should be only one npc
					speakerID = conn.RecNPCIDs[0]
					speaker = getNpcName(speakerID)
					spkDialog = getNpcDialog(speakerID)
				}

				if conn.Name == "ATTACK" {
					monsterIDs = append(monsterIDs, conn.RecMobIDs...)
				}
			}

			//testing
			// fmt.Println("| DEBUG end |")

			// add quit option
			options = append(options, "QUIT")

			fmt.Println("")
			choice := userinputs.MultiChoiceAnswer("What would you like to do?", options)

			if choice == "TALK" {
				// do talking things
				// clear screen
				callClear()

				// show dialog
				fmt.Println(speaker)
				fmt.Println(spkDialog)
				fmt.Println("")
				// set npc talked
				setNpcTalked(speakerID)
				// give items?
				npcInv := getNpcInv(speakerID)
				for _, itemID := range npcInv {
					player.InvIDs = append(player.InvIDs, itemID)
					fmt.Println("Player recieved " + getItemName(itemID))
				}

				// pause and wait for user
				_ = userinputs.RequestAnswer("1: Continue...")

				// set cell to EndCellID
				for i, op := range options {
					if op == choice {
						player.PLocID = player.CLocID
						player.CLocID = optionIds[i]
					}
				}
			} else if choice == "ATTACK" {
				// do attack things here
				for _, monID := range monsterIDs {
					if attackMob(monID) {
						// set mob dead
						setMobDead(monID)
						// get any items they may have had on them
						monInv := getMobInv(monID)
						for _, itemID := range monInv {
							player.InvIDs = append(player.InvIDs, itemID)
							fmt.Println("Player recieved " + getItemName(itemID))
						}
					} else {
						// set player dead
						fmt.Println("Player was slain by " + getMobName(monID))

						// pause and wait for user
						_ = userinputs.RequestAnswer("1: The End")

						// end game
						endGame = true
						return
					}
				}

				// pause and wait for user
				_ = userinputs.RequestAnswer("1: Continue...")

				// set cell to EndCellID
				for i, op := range options {
					if op == choice {
						player.PLocID = player.CLocID
						player.CLocID = optionIds[i]
					}
				}
			} else if choice == "END" || choice == "QUIT" {
				// end game
				endGame = true
				return
			} else {
				// movement commands
				// add any items found in current cell
				cellInv := getCellInv(player.CLocID)
				if len(cellInv) > 0 {
					// clear screen
					callClear()

					for _, itemID := range cellInv {
						player.InvIDs = append(player.InvIDs, itemID)
						fmt.Println("Player recieved " + getItemName(itemID))
					}

					// pause and wait for user
					_ = userinputs.RequestAnswer("1: Continue...")
				}

				// set cell to EndCellID
				for i, op := range options {
					if op == choice {
						player.PLocID = player.CLocID
						player.CLocID = optionIds[i]
					}
				}
			}
		}
	}
	return
}

func attackMob(mobID int) bool {
	battleLoss := false
	mobHP := getMobHP(mobID)
	mobAtt := getMobAtt(mobID)
	mobDef := getMobDef(mobID)
	playHP := getPlayersTotalHP()
	playAtt := getPlayersTotalAtt()
	playDef := getPlayersTotalDef()

	for !battleLoss {
		// clear screen
		callClear()
		fmt.Println("[Player: " + strconv.Itoa(playHP) + " | " + getMobName(mobID) + ": " + strconv.Itoa(mobHP) + "]")
		fmt.Println("")

		// player goes first
		attack := randimizer.IntRandimizer(1, (20 + playAtt))
		// on successful hit
		if attack > mobDef {
			mobHP = mobHP - (attack - mobDef) // diffrense of attack and mob's def
			fmt.Println("Player dealt " + strconv.Itoa(attack-mobDef) + " points of damage to " + getMobName(mobID))
		} else {
			fmt.Println("Player missed " + getMobName(mobID))
		}

		// mobs turn
		// first check mobs hp
		if mobHP > 0 {
			attack := randimizer.IntRandimizer(1, (20 + mobAtt))
			if attack > playDef {
				playHP = playHP - (attack - playDef)
				fmt.Println(getMobName(mobID) + " dealt " + strconv.Itoa(attack-playDef) + " points of damage to Player")
				fmt.Println("")
			} else {
				fmt.Println(getMobName(mobID) + " missed Player")
				fmt.Println("")
			}
		} else {
			fmt.Println(getMobName(mobID) + " is Dead")
			fmt.Println("")
			// pause and wait for user
			_ = userinputs.RequestAnswer("1: Continue...")
			return true
		}

		// check player's HP
		if playHP <= 0 {
			battleLoss = true
			break
		}

		// pause and wait for user
		_ = userinputs.RequestAnswer("1: Next Round...")

	}

	return false
}

func getCellInv(cellID int) []int {
	var itemInv []int
	// check each cell
	for _, cell := range world.Cells {
		// comaire cell with the one where are checking
		if cell.ID == cellID {
			// if they match return npc.Dialog
			for _, item := range cell.Items {
				itemInv = append(itemInv, item.ID)
			}
		}
	}
	// return items
	return itemInv
}
