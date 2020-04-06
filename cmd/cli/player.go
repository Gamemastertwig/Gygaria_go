package main

func getPlayersTotalHP() int {
	total := player.HP
	// check each item in player inv
	for _, invID := range player.InvIDs {
		// get items HP, and add to total
		total = total + getItemHP(invID)
	}
	return total
}

func getPlayersTotalAtt() int {
	total := player.Attack
	// check each item in player inv
	for _, invID := range player.InvIDs {
		// get items Attack, and add to total
		total = total + getItemAtt(invID)
	}
	return total
}

func getPlayersTotalDef() int {
	total := player.Defense
	// check each item in player inv
	for _, invID := range player.InvIDs {
		// get items Defense, and add to total
		total = total + getItemDef(invID)
	}
	return total
}
