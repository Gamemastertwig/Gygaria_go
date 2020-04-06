package main

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

func getItemHP(itemID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check items in each cell
		for _, cellItem := range cell.Items {
			// compaire item(s) with one we are looking for
			if cellItem.ID == itemID {
				// if match return cellItem.HP
				return cellItem.HP
			}
		}

		// check each mob in each cell (if not found in cell directly)
		for _, mob := range cell.Mobs {
			// check each item on the mobs
			for _, mobItem := range mob.Items {
				// compaire item(s) with one we are looking for
				if mobItem.ID == itemID {
					// if match return mobItem.HP
					return mobItem.HP
				}
			}
		}

		// check each npc in each cell (if not found in cell or mobs directly)
		for _, npc := range cell.NPCs {
			// check each item on the mobs
			for _, npcItem := range npc.Items {
				// compaire item(s) with one we are looking for
				if npcItem.ID == itemID {
					// if match return npcItem.HP
					return npcItem.HP
				}
			}
		}
	}
	// otherwise return blank
	return 0
}

func getItemAtt(itemID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check items in each cell
		for _, cellItem := range cell.Items {
			// compaire item(s) with one we are looking for
			if cellItem.ID == itemID {
				// if match return cellItem.Attack
				return cellItem.Attack
			}
		}

		// check each mob in each cell (if not found in cell directly)
		for _, mob := range cell.Mobs {
			// check each item on the mobs
			for _, mobItem := range mob.Items {
				// compaire item(s) with one we are looking for
				if mobItem.ID == itemID {
					// if match return mobItem.Attack
					return mobItem.Attack
				}
			}
		}

		// check each npc in each cell (if not found in cell or mobs directly)
		for _, npc := range cell.NPCs {
			// check each item on the mobs
			for _, npcItem := range npc.Items {
				// compaire item(s) with one we are looking for
				if npcItem.ID == itemID {
					// if match return npcItem.Attack
					return npcItem.Attack
				}
			}
		}
	}
	// otherwise return blank
	return 0
}

func getItemDef(itemID int) int {
	// check each cell
	for _, cell := range world.Cells {
		// check items in each cell
		for _, cellItem := range cell.Items {
			// compaire item(s) with one we are looking for
			if cellItem.ID == itemID {
				// if match return cellItem.Defense
				return cellItem.Defense
			}
		}

		// check each mob in each cell (if not found in cell directly)
		for _, mob := range cell.Mobs {
			// check each item on the mobs
			for _, mobItem := range mob.Items {
				// compaire item(s) with one we are looking for
				if mobItem.ID == itemID {
					// if match return mobItem.Defense
					return mobItem.Defense
				}
			}
		}

		// check each npc in each cell (if not found in cell or mobs directly)
		for _, npc := range cell.NPCs {
			// check each item on the mobs
			for _, npcItem := range npc.Items {
				// compaire item(s) with one we are looking for
				if npcItem.ID == itemID {
					// if match return npcItem.Defense
					return npcItem.Defense
				}
			}
		}
	}
	// otherwise return blank
	return 0
}
