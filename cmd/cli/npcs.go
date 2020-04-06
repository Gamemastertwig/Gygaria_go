package main

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

func getNpcDialog(npcID int) string {
	// check each cell
	for _, cell := range world.Cells {
		// check npc in each cell
		for _, npc := range cell.NPCs {
			// comaire npcs with the one where are checking
			if npc.ID == npcID {
				// if they match return npc.Dialog
				return npc.Dialog
			}
		}
	}
	// otherwise return blank
	return ""
}

func getNpcInv(npcID int) []int {
	var itemInv []int
	// check each cell
	for _, cell := range world.Cells {
		// check npc in each cell
		for _, npc := range cell.NPCs {
			// comaire npcs with the one where are checking
			if npc.ID == npcID {
				// if they match return npc.Dialog
				for _, item := range npc.Items {
					itemInv = append(itemInv, item.ID)
				}
			}
		}
	}
	// return items
	return itemInv
}

func setNpcTalked(npcID int) {
	// check each cell
	for i := range world.Cells {
		// check npc in each cell
		for j := range world.Cells[i].NPCs {
			// comaire npcs with the one where are checking
			if world.Cells[i].NPCs[j].ID == npcID {
				// if they match make npc talked
				world.Cells[i].NPCs[j].HasTalked = true
			}
		}
	}
}
