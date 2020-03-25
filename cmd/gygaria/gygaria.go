// Package gygaria is a package to pre-install the Gygaria TBA (Text Based Adventure).
// By pre-install means creating the json file for the "Unnamed" TBA engine.
// Will also act as a test bed for the TBA engine's functionality
package gygaria

import (
	"github.com/Gamemastertwig/Gygaria_go/pkg/storycell"
)

// CreateGygaria loads all variables for the world of gygaria
func CreateGygaria() {
	// code to pre-fill in all of gygaria

	// ****************
	// ** Items      **
	// ****************
	var items []storycell.Item

	// Potion of Youth
	potion := storycell.Item{Name: "Potion of Youth", Disc: "The flask holds a liquid that seems to shift colors, the longer you stare into it the more you swear you can hear children's laughter.", HP: 0, Attack: 0, Defense: 0}
	items = append(items, potion)

	// Seal of the Magi
	seal := storycell.Item{Name: "Seal of the Magi", Disc: "The seal is attached to a cord that one would wear over their neck. On one side there is a picture of a young man you do not reconize and on the other side are strange words you cannot read.", HP: 1000, Attack: 0, Defense: 0}
	items = append(items, seal)

	// Cloak of Almost Invisibility
	cloak := storycell.Item{Name: "Cloak of Almost Invisibility", Disc: "", HP: 0, Attack: 0, Defense: 0}
	items = append(items, cloak)

	// Death's Coin
	coin := storycell.Item{Name: "Death's Coin", Disc: "", HP: 0, Attack: 0, Defense: 0}
	items = append(items, coin)

	// Armor of Spider Silk
	armor := storycell.Item{Name: "Armor of Spider Silk", Disc: "", HP: 0, Attack: 0, Defense: 1000}
	items = append(items, armor)

	// Sword of Dragon's Breath
	sword := storycell.Item{Name: "Sword of Dragon's Breath", Disc: "", HP: 0, Attack: 1000, Defense: 0}
	items = append(items, sword)

	// ****************
	// ** NPCs       **
	// ****************
	var npcs []storycell.NPC

	// Old Man
	oldMan := storycell.NPC{Name: "Old Man", Dialog: "It's about time you got here! I have been waiting almost 2 weeks for your arrival. You need to hurry if you are to save the princess in the castle to the north which is sealed with magic, the fairy to the west may know more of how to break the seal. You will need to get better armor and a better weapon. I hear the forest to the east may have magical armor. I also hear the dragon to the south-west has a magical sword. Hurry now your adventure awaits you. (With that the old man vanishes.)"}
	npcs = append(npcs, oldMan)

	// Lady
	lady := storycell.NPC{Name: "Lady", Dialog: "So you see me? Then surely you will be the one to save us all. Your future holds great power but only you can unlock it. Your power is not that of a dragon. They say a cloak can be found in a bone yard to the north. Make haste sparrow for your fate is that of our kingdom's. (With that the lady in rages vanishes.)"}
	npcs = append(npcs, lady)

	// Death
	death := storycell.NPC{Name: "Death", Dialog: "I cannot let you pass until I get me coin back. A wealth man tricked me and toke it,but I curse him to an early death. He should be buried near town. I cannot leave here, get it for me and I will grant you free passage for life."}
	npcs = append(npcs, death)

	// Demon
	demon := storycell.NPC{Name: "Demon", Dialog: "You are not the one that called me here. I have this Potion of Youth you can have if you sell me your soul or defeat me in battle. Do we make a DEAL or FIGHT? "}
	npcs = append(npcs, demon)

	// Demon (Deal)
	demonDeal := storycell.NPC{Name: "Demon (Deal)", Dialog: "I do not have need of your soul as of yet, but I will return when I do... (With that the demon takes flight and vanishes.)"}
	demonDeal.Items = append(demonDeal.Items, potion)
	npcs = append(npcs, demonDeal)

	// Fairy (Before Potion)
	fairyBefore := storycell.NPC{Name: "Fairy (Quest)", Dialog: "I don't want to speak to you unless you have my Potion of Youth. That demon stole it from me at the crossroads."}
	npcs = append(npcs, fairyBefore)

	// Fairy (Potion)
	fairyPotion := storycell.NPC{Name: "Fairy (Potion)", Dialog: "Thanks for getting the potion. This potion is only for magical beings. If a human had drank this they would have turned into a baby or young child. Here is a Seal of the Magi, it will allow entry into the castle."}
	fairyPotion.Items = append(fairyPotion.Items, seal)
	npcs = append(npcs, fairyPotion)

	// Fairy (After Potion)
	fairyAfter := storycell.NPC{Name: "Fairy", Dialog: "I have nothing further to say to you."}
	npcs = append(npcs, fairyAfter)

	// Dragon (Riddle)

	// Wizard
	wizard := storycell.NPC{Name: "Wizard", Dialog: "How dare you enter my castle... and with those dirty boots! I made this castle from the ground up and do you show respect? No, you just tromp all around like you're the king. I cannot believe the nerve of you... I mean honestly. Ugh! I'm the most powerful wizard in the land and you are just a worm. I have killed stronger men then you in my sl... (Before he gets a chance to finish his monologue you take a good slice out of his arm.)"}
	npcs = append(npcs, wizard)

	// ****************
	// ** Mobs       **
	// ****************
	// var mobs []storycell.Mob

	// Orc 1
	// Orc 2
	// Cannibal 1
	// Cannibal 2
	// Cannibal 3
	// Cannibal 4
	// Cannibal 5
	// Wolf 1
	// Wolf 2
	// Wolf 3
	// Wolf 4
	// Sleep
	// Demon
	// Skeleton 1
	// Skeleton 2
	// Sandworm
	// Zombie
	// Giant Spider
	// Dragon
	// Wizard

	// ****************
	// ** MapCells   **
	// ****************
	// var cells []storycell.MapCell

	// Orc Camp
	// Orc Camp (Empty)
	// Castle Courtyard
	// Castle Couttyard (Seal)
	// Castle (Wizard Ending)
	// Land Slide
	// Dense Forest
	// Backwater Village
	// Backwater Village (Empty)
	// Wolves' Den
	// Wolves' Den (Empty)
	// Spring Meadow
	// Spring Meadow (Empty)
	// Crossroads
	// Crossroads (Deal)
	// Crossroads (Empty)
	// Bridge
	// Boneyard
	// Boneyard (Empty)
	// Boneyard (Thieves' Ending)
	// Fairy Glade
	// Fairy Glade (Quest)
	// Fairy Glade (Potion)
	// Crystal Lake
	// Crystal Lake (Quest)
	// Caravan
	// Plains
	// Desert
	// Desert (Empty)
	// Grave Keeper's Hovel
	// Grave Keeper's Hovel (Deed Ending)
	// Town
	// Town (Old Man)
	// Town (Lady in Rags)
	// Farm
	// Spider Forest
	// Spider Forest (Empty)
	// Mountain
	// Dragon's Lair
	// Dragon's Lair (Riddle)
	// Dragon's Lair (Sneak)
	// Dragon's Lair (Empty)
	// Docks
	// Docks (Sailor's Ending)
	// Ocean
	// Cliff

}
