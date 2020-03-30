package main

import (
	"fmt"

	"github.com/Gamemastertwig/Gygaria_go/pkg/userinputs"

	"github.com/Gamemastertwig/Gygaria_go/cmd/gygaria"
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
)

var world tbastory.World
var player tbastory.Player

func main() {
	world = gygaria.CreateGygaria()
	player = gygaria.CreatePlayer()
	displayCurrentCell()
}

func displayCurrentCell() {
	currentCellID := player.CLocID

	for _, cell := range world.Cells {
		if cell.ID == currentCellID {
			fmt.Println(cell.Title)
			fmt.Println(cell.Disc)
			var options []string
			var optionIds []int
			for _, conn := range cell.Conns {
				connGood := false
				pass := 0
				count := 0
				if conn.NeedItem {
					// check inv for items in conn
					for _, itemID := range conn.RecItemIDs {
						count++
						for _, invID := range player.InvIDs {
							if invID == itemID {
								pass++
							}
						}
					}
				}
				if conn.NeedNotItem {
					// check inv for the absense of items in conn
					for _, itemID := range conn.RecItemIDs {
						count++
						for _, invID := range player.InvIDs {
							if invID != itemID {
								pass++
							}
						}
					}
				}
				if conn.NeedMobAlive {
					// check that mob is alive for mobs in conn
					for _, recMobID := range conn.RecMobIDs {
						count++
						for _, cell := range world.Cells {
							for _, mob := range cell.Mobs {
								if mob.ID == recMobID {
									if !mob.IsDead {
										pass++
									}
								}
							}
						}
					}
				}
				if conn.NeedMobKilled {
					// check that mob is dead for mobs in conn
					for _, recMobID := range conn.RecMobIDs {
						count++
						for _, cell := range world.Cells {
							for _, mob := range cell.Mobs {
								if mob.ID == recMobID {
									if mob.IsDead {
										pass++
									}
								}
							}
						}
					}
				}
				if conn.NeedNPCTalk {
					// check that npc talked for npcs in conn
					for _, recNpcID := range conn.RecNPCIDs {
						count++
						for _, cell := range world.Cells {
							for _, npc := range cell.NPCs {
								if npc.ID == recNpcID {
									if npc.HasTalked {
										pass++
									}
								}
							}
						}
					}
				}
				if conn.NeedNPCNoTalk {
					// check that npc have not talked for npcs in conn
					for _, recNpcID := range conn.RecNPCIDs {
						count++
						for _, cell := range world.Cells {
							for _, npc := range cell.NPCs {
								if npc.ID == recNpcID {
									if !npc.HasTalked {
										pass++
									}
								}
							}
						}
					}
				}
				fmt.Printf("Count: %v ", count)
				fmt.Printf("Pass: %v ", pass)
				if pass == count {
					connGood = true
				}
				if connGood {
					options = append(options, conn.Name)
					optionIds = append(optionIds, conn.EndCellID)
					//fmt.Println(conn.Name)
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
