package main

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

func getMobHP(mobID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check mob(s) in each cell
		for _, mob := range cell.Mobs {
			// comaire mob(s) with the one we are checking
			if mob.ID == mobID {
				// if they are the same return mob.HP
				return mob.HP
			}
		}
	}
	// otherwise return blank
	return 0
}

func getMobAtt(mobID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check mob(s) in each cell
		for _, mob := range cell.Mobs {
			// comaire mob(s) with the one we are checking
			if mob.ID == mobID {
				// if they are the same return mob.Attack
				return mob.Attack
			}
		}
	}
	// otherwise return blank
	return 0
}

func getMobDef(mobID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check mob(s) in each cell
		for _, mob := range cell.Mobs {
			// comaire mob(s) with the one we are checking
			if mob.ID == mobID {
				// if they are the same return mob.Defense
				return mob.Defense
			}
		}
	}
	// otherwise return blank
	return 0
}

func getMobInv(mobID int) []int {
	var itemInv []int
	// check each cell
	for _, cell := range world.Cells {
		// check npc in each cell
		for _, mob := range cell.Mobs {
			// comaire npcs with the one where are checking
			if mob.ID == mobID {
				// if they match return npc.Dialog
				for _, item := range mob.Items {
					itemInv = append(itemInv, item.ID)
				}
			}
		}
	}
	// return items
	return itemInv
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

func setMobDead(mobID int) {
	// check each cell
	for i := range world.Cells {
		// check npc in each cell
		for j := range world.Cells[i].Mobs {
			// comaire npcs with the one where are checking
			if world.Cells[i].Mobs[j].ID == mobID {
				// if they match make mob dead
				world.Cells[i].Mobs[j].IsDead = true
			}
		}
	}
}
