// Package gygaria is a package to pre-install the Gygaria TBA (Text Based Adventure).
// By pre-install means creating the json file for the "Unnamed" TBA engine.
// Will also act as a test bed for the TBA engine's functionality
package gygaria

import (
	"github.com/Gamemastertwig/Gygaria_go/pkg/tbastory"
)

// CreateGygaria loads all variables for the world of gygaria
func CreateGygaria() tbastory.World {
	// Data to pre-fill in all of gygaria

	// ***********
	// ** World **
	// ***********

	// Gygaria
	var gygaria tbastory.World

	// ***********
	// ** Items **
	// ***********

	// Potion of Youth
	potion := tbastory.Item{ID: 1, Name: "Potion of Youth", Disc: "The flask holds a liquid that seems to shift colors, the longer you stare into it the more you swear you can hear children's laughter.", HP: 0, Attack: 0, Defense: 0}

	// Seal of the Magi
	seal := tbastory.Item{ID: 2, Name: "Seal of the Magi", Disc: "The seal is attached to a cord that one would wear over their neck. On one side there is a picture of a young man you do not reconize and on the other side are strange words you cannot read.", HP: 1000, Attack: 0, Defense: 0}

	// Cloak of Almost Invisibility
	cloak := tbastory.Item{ID: 3, Name: "Cloak of Almost Invisibility", Disc: "The cloak seems to reflect light in a manner that makes the wear seem invisible. Quick movements or constant observation from an onlooker makes it apperent someone or something is present, hense the 'almost'.", HP: 1000, Attack: 0, Defense: 0}

	// Death's Coin
	coin := tbastory.Item{ID: 4, Name: "Death's Coin", Disc: "", HP: 0, Attack: 0, Defense: 0}

	// Armor of Spider Silk
	armor := tbastory.Item{ID: 5, Name: "Armor of Spider Silk", Disc: "", HP: 0, Attack: 0, Defense: 1000}

	// Sword of Dragon's Breath
	sword := tbastory.Item{ID: 6, Name: "Sword of Dragon's Breath", Disc: "", HP: 0, Attack: 1000, Defense: 0}

	// **********
	// ** NPCs **
	// **********

	// Old Man
	oldMan := tbastory.NPC{ID: 1, Name: "Old Man", Dialog: "It's about time you got here! I have been waiting almost 2 weeks for your arrival. You need to hurry if you are to save the princess in the castle to the north which is sealed with magic, the fairy to the west may know more of how to break the seal. You will need to get better armor and a better weapon. I hear the forest to the east may have magical armor. I also hear the dragon to the south-west has a magical sword. Hurry now your adventure awaits you. (With that the old man vanishes.)"}

	// Lady
	lady := tbastory.NPC{ID: 2, Name: "Lady", Dialog: "So you see me? Then surely you will be the one to save us all. Your future holds great power but only you can unlock it. Your power is not that of a dragon. They say a cloak can be found in a bone yard to the north. Make haste sparrow for your fate is that of our kingdom's. (With that the lady in rages vanishes.)"}

	// Death
	death := tbastory.NPC{ID: 3, Name: "Death", Dialog: "I cannot let you pass until I get me coin back. A wealth man tricked me and toke it,but I curse him to an early death. He should be buried near town. I cannot leave here, get it for me and I will grant you free passage for life."}

	// Demon
	demon := tbastory.NPC{ID: 4, Name: "Demon", Dialog: "You are not the one that called me here. I have this Potion of Youth you can have if you sell me your soul or defeat me in battle. Do we make a DEAL or FIGHT? "}

	// Demon (Deal)
	demonDeal := tbastory.NPC{ID: 5, Name: "Demon (Deal)", Dialog: "I do not have need of your soul as of yet, but I will return when I do... (With that the demon takes flight and vanishes.)"}
	demonDeal.Items = append(demonDeal.Items, potion)

	// Fairy
	fairy := tbastory.NPC{ID: 6, Name: "Fairy", Dialog: "I have nothing to say to you."}

	// Fairy (Before Potion)
	fairyBefore := tbastory.NPC{ID: 7, Name: "Fairy (Quest)", Dialog: "I don't want to speak to you unless you have my Potion of Youth. That demon stole it from me at the crossroads."}

	// Fairy (Potion)
	fairyPotion := tbastory.NPC{ID: 8, Name: "Fairy (Potion)", Dialog: "Thanks for getting the potion. This potion is only for magical beings. If a human had drank this they would have turned into a baby or young child. Here is a Seal of the Magi, it will allow entry into the castle."}
	fairyPotion.Items = append(fairyPotion.Items, seal)

	// Dragon (Riddle)
	dragon := tbastory.NPC{ID: 9, Name: "Dragon", Dialog: ""}
	dragon.Items = append(dragon.Items, sword)

	// **********
	// ** Mobs **
	// **********

	// Orcs
	orc1 := tbastory.Mob{ID: 1, Name: "Kargrosh", Disc: "Orc", HP: 0, Attack: 0, Defense: 0}
	orc2 := tbastory.Mob{ID: 2, Name: "Kromeash", Disc: "Orc", HP: 0, Attack: 0, Defense: 0}

	// Cannibals
	cannibal1 := tbastory.Mob{ID: 3, Name: "Warib Gest", Disc: "Cannibal", HP: 0, Attack: 0, Defense: 0}
	cannibal2 := tbastory.Mob{ID: 4, Name: "Deffery Jahmer", Disc: "Cannibal", HP: 0, Attack: 0, Defense: 0}
	cannibal3 := tbastory.Mob{ID: 5, Name: "Mtela Saris", Disc: "Cannibal", HP: 0, Attack: 0, Defense: 0}
	cannibal4 := tbastory.Mob{ID: 6, Name: "Filbert Ash", Disc: "Cannibal", HP: 0, Attack: 0, Defense: 0}
	cannibal5 := tbastory.Mob{ID: 7, Name: "Mermin Aiwes", Disc: "Cannibal", HP: 0, Attack: 0, Defense: 0}

	// Wolves
	wolf1 := tbastory.Mob{ID: 8, Name: "Gray", Disc: "Wolf", HP: 0, Attack: 0, Defense: 0}
	wolf2 := tbastory.Mob{ID: 9, Name: "Baga", Disc: "Wolf", HP: 0, Attack: 0, Defense: 0}
	wolf3 := tbastory.Mob{ID: 10, Name: "Balltoe", Disc: "Wolf", HP: 0, Attack: 0, Defense: 0}
	wolf4 := tbastory.Mob{ID: 11, Name: "Lowkey", Disc: "Wolf", HP: 0, Attack: 0, Defense: 0}

	// Sleep
	sleep := tbastory.Mob{ID: 25, Name: "Slumber", Disc: "Unnatural Sleep", HP: 0, Attack: 0, Defense: 0}

	// Demon
	demonAttack := tbastory.Mob{ID: 27, Name: "Harry", Disc: "Demon", HP: 0, Attack: 0, Defense: 0}
	demonAttack.Items = append(demonAttack.Items, potion)

	// Skeletons
	skele1 := tbastory.Mob{ID: 28, Name: "Jaeke", Disc: "Skeleton", HP: 0, Attack: 0, Defense: 0}
	skele2 := tbastory.Mob{ID: 15, Name: "Le-bone-ski", Disc: "Skeleton", HP: 0, Attack: 0, Defense: 0}
	skele2.Items = append(skele2.Items, cloak)

	// Sandworm
	sandworm := tbastory.Mob{ID: 16, Name: "Sandy", Disc: "Sandworm", HP: 0, Attack: 0, Defense: 0}

	// Zombie
	zombie := tbastory.Mob{ID: 29, Name: "Stiff McBrains", Disc: "Zombie", HP: 0, Attack: 0, Defense: 0}
	zombie.Items = append(zombie.Items, coin)

	// Giant Spider
	spider := tbastory.Mob{ID: 30, Name: "Arragogy", Disc: "Giant Spider", HP: 0, Attack: 0, Defense: 0}
	spider.Items = append(spider.Items, armor)

	// Dragon
	dragonAttack := tbastory.Mob{ID: 31, Name: "Dr. Agon", Disc: "Dragon", HP: 0, Attack: 0, Defense: 0}
	dragonAttack.Items = append(dragonAttack.Items, sword)

	// Wizard
	wizardAttack := tbastory.Mob{ID: 20, Name: "The Mad Mastermind", Disc: "Wizard", HP: 0, Attack: 0, Defense: 0}

	// *******************
	// ** Cells (Start) **
	// *******************
	// Each cell must have a unique ID

	// Orc Camp
	orcCamp := tbastory.MapCell{ID: 1, Title: "Orc Camp", Disc: "You see many small huts that are shabbily made. Two hulking figures stand about the place. They do not seem to be happy that you are in their camp and begine murmuring something at you in a guttural language you don’t understand. Then they come towards you with swords raised."}
	orcCamp.Mobs = append(orcCamp.Mobs, orc1)
	orcCamp.Mobs = append(orcCamp.Mobs, orc2)

	// Orc Camp (Empty)
	orcCampEmpty := tbastory.MapCell{ID: 2, Title: "Orc Camp (Empty)", Disc: "You see many small huts that are shabbily made. They are just a bit more 'shably' after your fight with the orcs that used to live here."}

	// Castle Courtyard
	castleCourtyard := tbastory.MapCell{ID: 3, Title: "", Disc: "As you progress onward, you see that the small castle. There is a large monument with a glowing blue center. As you get closer to the castle, there is a large glowing barrier where the door is at as well. It blocks you from entering the castle."}

	// Castle Courtyard (Seal)
	castleCourtyardSeal := tbastory.MapCell{ID: 4, Title: "", Disc: "As you progress onward, you see that the small castle. There is a large monument with a glowing blue center. As you get closer to the castle, there is a large glowing barrier where the door is at as well. When the seal is brought close the runes begin to fade and you can OPEN the door."}

	// Castle
	castle := tbastory.MapCell{ID: 5, Title: "Castle", Disc: "You are now inside the castle having completed most of your quest. You have beaten a dragon, talked with a demon, slain a zombie and cannibals as well. So many foes have vanquished before you, but so close to the end, you know your most difficult battle is before you. In this royal palace, it is covered in luxury. With large statues and comfortable thrones. A young wizard sits on the largest throne, as you make eye contact he begins to speak. How dare you enter my castle... and with those dirty boots! I made this castle from the ground up and do you show respect? No, you just tromp all around like you're the king. I cannot believe the nerve of you... I mean honestly. Ugh! I'm the most powerful wizard in the land and you are just a worm. I have killed stronger men then you in my sl... (Before he gets a chance to finish his monologue you take a good slice out of his arm.)"}
	castle.Mobs = append(castle.Mobs, wizardAttack)

	// Castle (Wizard Ending)
	castleEnd := tbastory.MapCell{ID: 6, Title: "The End", Disc: "As you strike the final blow the wizard falls to the ground in a whimper. 'We could have adopted beautiful children' he coughs. You turn to see a chunky little mushroom headed man in a cage. As you open the cage door he remarks to you that your princess is in another castle. You punch him in his big spotted head and continue your life until one day someone says that a giant turtle kidnapped their princess and toke her to a pipe world. A hero's work is never done..."}

	// Land Slide
	landSlide := tbastory.MapCell{ID: 7, Title: "Land Slide", Disc: "Whatever was in this place has long worn away from the sands of time. A rock slide now blocks the opening to the north that leads up a large cliff."}

	// Dense Forest
	forest := tbastory.MapCell{ID: 8, Title: "Dense Forest", Disc: "As you wander through the forest, you start to become disoriented. You try to walk through the forest but you find yourself on the outskirts in the exact place you came in."}

	// Backwater Village
	village := tbastory.MapCell{ID: 9, Title: "Backwater Village", Disc: "As you enter the town you notice that buildings are under-maintained, but the building structure leads you to believe that at one time a beautiful place. The people of the village seem to be huddled around a large fire. You move closer to investigate. You see that the people are poorly dressed and their clothes are raggedy or worse. Their bodies are unshaven, their hair is matted and they are eating something. They are all tearing at a piece of meat and it looks to be human... Suddenly they all turn to face you, reaching for their primitive weapons and start licking their lips in anticipation."}
	village.Mobs = append(village.Mobs, cannibal1)
	village.Mobs = append(village.Mobs, cannibal2)
	village.Mobs = append(village.Mobs, cannibal3)
	village.Mobs = append(village.Mobs, cannibal4)
	village.Mobs = append(village.Mobs, cannibal5)

	// Backwater Village (Empty)
	villageEmpty := tbastory.MapCell{ID: 10, Title: "Backwater Village (Empty)", Disc: "As you enter the town you notice that buildings are under-maintained, but the building structure leads you to believe that at one time a beautiful place."}

	// Wolves' Den
	den := tbastory.MapCell{ID: 11, Title: "Wolves' Den", Disc: "As you continue on your way you see a large cave with an opening too small for you to enter. Outside you can see a group of 4 wolves. You don't know how many wolves are there are in total but you know how many are coming your way. "}
	den.Mobs = append(den.Mobs, wolf1)
	den.Mobs = append(den.Mobs, wolf2)
	den.Mobs = append(den.Mobs, wolf3)
	den.Mobs = append(den.Mobs, wolf4)

	// Wolves' Den (Empty)
	denEmpty := tbastory.MapCell{ID: 12, Title: "Wolves' Den (Empty)", Disc: "As you continue on your way you see a large cave with an opening too small for you to enter. You see evidance that wolves once lived here."}

	// Spring Meadow
	meadow := tbastory.MapCell{ID: 13, Title: "Spring Meadow", Disc: "As you walk the scenery changes from a beaten path into a soft spring meadow. There are birds chirping and flowers everywhere. You hadn't noticed it at first but the climate has slowly been changing to the perfect temperature."}

	// Spring Meadow (Seal)
	meadowSeal := tbastory.MapCell{ID: 14, Title: "Spring Meadow (Seal)", Disc: "Suddenly the seal in your pocket begins to become warm. Every foot step is becoming more difficult. Suddenly a voice breaks out over the air. The voice sings of a soft bed of grass where you can lay your head'. You glance to your right and see a large group of fresh baby clovers. You must FIGHT your urge to slumber."}
	meadowSeal.Mobs = append(meadowSeal.Mobs, sleep)

	// Spring Meadow (Empty)
	meadowEmpty := tbastory.MapCell{ID: 15, Title: "Spring Meadow (Empty)", Disc: "As you walk the scenery changes from a beaten path into a soft spring meadow. The seal begins to warm again, so it does not appear as calming as before as you recall your battle with sleep from before. However, nothing happens."}

	// Crossroads
	crossroads := tbastory.MapCell{ID: 16, Title: "Crossroads", Disc: "You come to a cross roads. Off to side of the road you see a mound of dirt a child has been playing in. A large figure appears before you. It is big terrible and smells like sulfur. You had heard of demons but had assumed they were just a myth. His big bat-like wings beat through the air, but he doesn't seem to be attacking you. Maybe you can TALK to him or ATTACK first."}
	crossroads.Mobs = append(crossroads.Mobs, demonAttack)
	crossroads.NPCs = append(crossroads.NPCs, demon)

	// Crossroads (Demon)
	crossroadsDeal := tbastory.MapCell{ID: 17, Title: "Crossroads (Demon)", Disc: "You come to a cross roads. The large demon still sits here waiting for your response. Take the DEAL or you could ATTACK and take the potion from his dead corpse."}
	crossroadsDeal.Mobs = append(crossroadsDeal.Mobs, demonAttack)
	crossroadsDeal.NPCs = append(crossroadsDeal.NPCs, demonDeal)

	// Crossroads (Empty)
	crossroadsEmpty := tbastory.MapCell{ID: 18, Title: "Crossroads (Empty)", Disc: "You come to a cross roads. Off to side of the road you see a mound of dirt a child has been playing in. There are no other signs of life."}

	// Bridge
	bridge := tbastory.MapCell{ID: 19, Title: "Bridge", Disc: "The road leads you to a bridge. The bridge is old and well worn. It's the rickety kind that you can feel sway with every step. You are wary of it because you cannot swim. But even that hedgehog guy does water zones and he cannot swim either."}

	// Boneyard
	boneyard := tbastory.MapCell{ID: 20, Title: "Boneyard", Disc: "Discription: Let's be honest no one wants to be in a bone yard. It's a typical graveyard. Well except for the graves are open and there are skulls on pikes. Ok, so it is a weird graveyard. Just outside the bone yard is a band of what appear to be thieves. Once more all of them are women, and it appears they wish your accompaniment. If you JOIN them you could have riches the likes of which you have never seen... at the price of your princess. You see one grave that is nearly untouched. Upon closer inspection is has a cloak laid over it. Do you GRAB it?"}

	// Boneyard (Grab)
	boneyardGrab := tbastory.MapCell{ID: 21, Title: "Boneyard (Grab)", Disc: "As you reach for the cloak you hear your wrist pop, but wait. That wasn’t your wrist! The grave soil moves and a cpulr skeletons start to emerging covered in ruins. You have heard of people protecting their graves in this manner, but never actually seen it. At any rate you have made them angry and now they have a bone to pick with you, or off you. The battle ensues..."}
	boneyardGrab.Mobs = append(boneyardGrab.Mobs, skele1)
	boneyardGrab.Mobs = append(boneyardGrab.Mobs, skele2)

	// Boneyard (Empty)
	boneyardEmpty := tbastory.MapCell{ID: 22, Title: "Boneyard (Empty)", Disc: "Let's be honest no one wants to be in a bone yard. It's a typical graveyard. Well except for the graves are open and there are skulls on pikes. Ok, so it is a weird graveyard.Just outside the bone yard is a band of what appear to be thieves. Once more all of them are women, and it appears they wish your accompaniment. If you JOIN them you could have riches the likes of which you have never seen... at the price of your princess."}

	// Boneyard (Thieves' Ending)
	boneyardEnd := tbastory.MapCell{ID: 23, Title: "The End", Disc: "As you approach the girls begin to talk in a foreign language and flash you a smile. They take you to their ship, and pull you aboard. You smile at the fact of how they make over you and you travel on the sea for many months. When you arrive at shore you are quickly shackled and thrown into a cage the women laugh and sail away. Shortly after you are taken to a market where many foreigners poke and prod at you. You are taken to a large house where you are forced to work as a slave. You live in a hole in the ground for many months and beaten regularly. One day you try to run away but being that you don’t know the area you are quickly found and punished. You return to the cage minus a few fingers and aren't given any medical treatment. A few months later you die from a combination of gangrene, black plague, and malnourishment."}

	// Fairy Glade
	glade := tbastory.MapCell{ID: 24, Title: "Fairy Glade", Disc: "As you depart from the boat you see a field covered in clovers and woodland creatures. "}
	glade.NPCs = append(glade.NPCs, fairy)

	// Fairy Glade (Quest)
	gladeQuest := tbastory.MapCell{ID: 25, Title: "Fairy Glade (Quest)", Disc: "As you depart from the boat you see a field covered in clovers and woodland creatures. your eyes are drawn to a glowing fairy. You feel earily compeled to TALK to her"}
	gladeQuest.NPCs = append(gladeQuest.NPCs, fairyBefore)

	// Fairy Glade (Potion)
	gladePotion := tbastory.MapCell{ID: 26, Title: "Fairy Glade (Potion)", Disc: "As you depart from the boat you see a field covered in clovers and woodland creatures. your eyes are drawn to a glowing fairy. She beconds you over to TALK."}
	gladePotion.NPCs = append(gladePotion.NPCs, fairyPotion)

	// Crystal Lake
	lake := tbastory.MapCell{ID: 27, Title: "Crystal Lake", Disc: "You come to Crystal Lake. Although you cannot recall ever being here before, you realize why it is called Crystal Lake. It's the most reflective body of water you have ever seen. A small river boat is docked on the lake. "}

	// Crystal Lake (Quest)
	lakeQuest := tbastory.MapCell{ID: 28, Title: "Crystal Lake (Quest)", Disc: "You come to Crystal Lake. Although you cannot recall ever being here before, you realize why it is called Crystal Lake. It's the most reflective body of water you have ever seen. A small river boat is docked on the lake. Although it is day time there is a man is shrouded in darkness. The man is wearing a cloak and his figure is trim to a sickly amount. His face is hidden by the cloak but you are almost afraid to look at it. An empty amulet rests on his chest with the word 'death cometh'. At one point it seems to have housed a coin. You must TALK to him to get passage over the lake."}
	lakeQuest.NPCs = append(lakeQuest.NPCs, death)

	// Caravan
	caravan := tbastory.MapCell{ID: 29, Title: "Caravan", Disc: "If you had thought town was empty before, this was the reason. As you progress down the road you find it impossible not to stare at the gypsies entertaining the throngs of people. Not only do they entertain, but many have stalls where they are selling 'magical' or 'enchanted' items. One particular item that catches your interest is a music box that will grant wish if played under a full moon, but for only ten gold coins you think it must be a fake. Deciding this place holds nothing of relevance you continue on your way."}

	// Plains
	plains := tbastory.MapCell{ID: 30, Title: "Plains", Disc: "As you step onto the plains a gust of wind blows through the grass rustling it. You remember when you were a child your father to take you to such a place. As you ponder such things a duck shows itself from the brush. You attack but miss startling the mallard and it takes to the air far above your reach. A brown hound dog rises over the grass and chuckles at your misfortune... wait... WHAT THE HECK!?! The turn of events reminds you of your quest and you decide it best not to stand about doing nothing."}

	// Desert
	desert := tbastory.MapCell{ID: 31, Title: "Desert", Disc: "The grass you had been walking on begins to give way more and more until it is completely gone. In its place is course heat warmed sand. You begin to wonder if it is smart to find yourself here, after this entire place has no significance to your journey. Suddenly there is shifting of the sand below you. As you look down your gaze is met by the thousand teeth of the sand worm! Now is not the time for hesitation... RUN? Or FIGHT?"}
	desert.Mobs = append(desert.Mobs, sandworm)

	// Desert (Empty)
	desertEmpty := tbastory.MapCell{ID: 32, Title: "Desert (Empty)", Disc: "The grass you had been walking on begins to give way more and more until it is completely gone. In its place is course heat warmed sand. You begin to wonder if it is smart to find yourself here, after this entire place has no significance to your journey."}

	// Grave Keeper's Hovel
	hovel := tbastory.MapCell{ID: 33, Title: "Grave Keeper's Hovel", Disc: "The grave keeper isn't the hunch back living in a belfry that you had expected. In fact he looks familiar. You're almost sure you know him. As you approach he calls your name. Apparently he was an old friend of your dads. He tells you that your dad saved his life and he wants to repay him by giving you some nice land he inherited. You could TAKE the offer, move on, have a family and move into the roomy estate. But wouldn't you always regret letting down your princess?"}

	// Grave Keeper's Hovel (Deed Ending)
	hovelEnd := tbastory.MapCell{ID: 34, Title: "The End", Disc: "You move to the country and live in a manor on the outskirts of town. Soon after, you meet a spouse and fall in love. Being that you are a young couple a few years later you have children, two twins. You live in peace for a few years and even become popular among the town. One day someone begins leaving threats of death on your door. Your spouse is terrified and asks you to move them out immediately. Reluctantly you begin to try and pack. About a week later as you are packing a man comes down the stairs with your children and hands you a knife and tells you to kill your children. You refuse and he tells you he will torture your entire family and that this is giving them an easy way out. Reluctantly you take the knife and try to end your kids' life. Unable to perform the deed you quickly turn the knife on yourself. As your life force seeps from your body, you think of how it may have been different if you had stayed and helped the princess."}

	// Town Graveyard
	graveyard := tbastory.MapCell{ID: 35, Title: "Town Graveyard", Disc: "It's a town grave yard... what do you want me to say? The graves are all nicely kept and there is a large mausoleum in the center (probably were the rich are laid to rest). You also feel a need to INSPECT the mausoleum."}

	// Town Graveyard (Inspect)
	graveyardInspect := tbastory.MapCell{ID: 36, Title: "Town Graveyard (Inspect)", Disc: "The mausoleum door is unlocked. You are surprised at the cleanliness of the place. There in the middle of the family insignia is Death's coin. 'That was easy.' you smirk. As you turn you are greeted by an shambling figure. The zombie convulses and attacks without warning."}
	graveyardInspect.Mobs = append(graveyardInspect.Mobs, zombie)

	// Town Graveyard (Empty)
	graveyardEmpty := tbastory.MapCell{ID: 37, Title: "Town Graveyard (Empty)", Disc: "It's a town grave yard... what do you want me to say? The graves are all nicely kept and there is a large mausoleum in the center (probably were the rich are laid to rest)."}

	// Town
	town := tbastory.MapCell{ID: 38, Title: "Town", Disc: "You are in the town square. What few people are here are busy doing there every day jobs and seem not to notice you."}

	// Town (Old Man)
	townMan := tbastory.MapCell{ID: 39, Title: "Town (Old Man)", Disc: "You are in the town square. What few people are here are busy doing there every day jobs and seem not to notice you. An old man dressed in purple robes keeps waving at you. It looks like he wants to TALK."}
	townMan.NPCs = append(townMan.NPCs, oldMan)

	// Town (Lady in Rags)
	townLady := tbastory.MapCell{ID: 40, Title: "Town (Lady in Rags)", Disc: "You are in the town square. What few people are here are busy doing there every day jobs and seem not to notice you. You notice a lady in rags. You are sure she was not here before. She keeps staring at you maybe she wants to TALK."}
	townLady.NPCs = append(townLady.NPCs, lady)

	// Farm
	farm := tbastory.MapCell{ID: 41, Title: "Farm", Disc: "You arrive on a field. The corn and beans grow to about mid thigh, being as it is midsummer and the cattle graze in their pins. A small house rest on a hill near-by beyond that there appears to be nothing of interest. Upon further inspection a farmer is diligently at work tending to his crop. His gruff appearance and the way he's been eyeing you leads you to believe he's not much for talking."}

	// Spider Forest
	spiderForest := tbastory.MapCell{ID: 42, Title: "Spider Forest", Disc: "Upon entering the forest you begin to get an eerie feeling. Looking around you, you notice the trees seem to get closer and closer. You also notice many webs which appear to be made of metallic silk. In some webs in front of you is a suite armor. As you approach the web a thick string of saliva drops just inches from your foot, bubbling at the  grass in front of you dissolving it straight to the soil. Looking above you see the figure of a large hairy eight legged being. As it encroaches upon you see a gargantuan spider."}
	spiderForest.Mobs = append(spiderForest.Mobs, spider)

	// Spider Forest (Empty)
	spiderForestEmpty := tbastory.MapCell{ID: 43, Title: "Spider Forest (Empty)", Disc: "Upon entering the forest you begin to get an eerie feeling. Looking around you, you notice the trees seem to get closer and closer. You also notice many webs which appear to be made of metallic silk. Nothing seems important here anymore and it appears you can only turn back."}

	// Mountain
	mountain := tbastory.MapCell{ID: 44, Title: "Mountains", Disc: "As you have travelled past the Grave Keepers home, you find yourself at the base of a mountain. Surely you can climb the mountain, but why? There is no defined path and the trip could take days, best to head back the way you came."}

	// Dragon's Lair
	dragonLair := tbastory.MapCell{ID: 45, Title: "Dragon's Lair", Disc: "You have come to a cave. It is a lot warmer than you would have expected from a cave this size and being located this close to the ocean. Smell of decay is heavy in the air and ground is littered with mangled corpses. Oddly the bones are bare although most have large incisions from what appears to be very large teeth. On the far side of the cave is the slumbering mass of a large black dragon. Behind it lays a claymore lodged in gold. The hilt is golden and inlayed with precious gems. You curse you luck knowing that even in this light there is no way to get passed him with what you have. If you had the fabled 'Cloak of almost Invisibility' you may just be able to SNEAK past. Why could it not be a giant turtle? You may have to FIGHT your way out of this mess. You hear dragon's will sometime ask riddles. Maybe you can TALK to the dragon to get access to what is on the other side. Hopefully you are smarter than the dragon. Otherwise you may end up as dinner."}
	dragonLair.Mobs = append(dragonLair.Mobs, dragonAttack)

	// Dragon's Lair (Riddle)
	dragonLairRiddle := tbastory.MapCell{ID: 46, Title: "Dragon's Lair (Riddle)", Disc: "You approch the dragon with hands outright. The dragon looks at you lazely and laughs. I could easily make you my snack, but you ammuse me. Answer me this riddle and you are free to take one item from my hord. Fail to answer correctly and a snack I will have. It is so strong, it does not break, It is so powerful, it penetrates a lake, Somrtimes it is weak as a twig left out in the sun, It is so wonderful it gives us the power to have fun. What is it? A: Darkness B: Light C: Moon D: Sun"}
	dragonLairRiddle.Items = append(dragonLairRiddle.Items, sword)

	// Dragon's Lair (Sneak)
	dragonLairSneak := tbastory.MapCell{ID: 47, Title: "Dragon's Lair (Sneak)", Disc: "You place on the Cloak of Almost Invisibility and begin to fade. You do not fade completely but with the lower light of the cave you are a lot hard to see. As you creep past the dragon, it begins to move then look around it stares in your direction for what seems like forever. Finally it seems content that no one is in the cave with it and it goes back to sleep. You continue your sneak around and get the sword. You experience no troubles on your return trip."}
	dragonLairSneak.Items = append(dragonLairSneak.Items, sword)

	// Dragon's Lair (Empty)
	dragonLairEmpty := tbastory.MapCell{ID: 48, Title: "Dragon's Lair (Empty)", Disc: "You have come to a cave. It is a lot warmer than you would have expected from a cave this size and being located this close to the ocean. Smell of decay is heavy in the air and ground is littered with mangled corpses. Oddly the bones are bare although most have large incisions from what appears to be very large teeth. There is nothing left of interest here, you should probably turn back."}

	// Docks
	docks := tbastory.MapCell{ID: 49, Title: "Docks", Disc: "As you approach the docks you see a young man who is frantic. Upon inquiry he tells you he is short a crew man and can't make his voyage without the extra man power. He offers you a healthy sum of gold to join his expedition. In doing so you could effectively disregard your troubles in finding the princess and become a seaman. Besides heroes are overrated... or are they? You could JOIN the captain for a life on the sea."}

	// Docks (Sailor's Ending)
	docksEnd := tbastory.MapCell{ID: 50, Title: "The End", Disc: "You climb on the boat and become comfortable as a sailor. It turns out it was a fishing boat and you make a plentiful haul for a few years. One year you are informed of a rare group of Gygarian trout in the northern oceans. Your crew travels north and after many months they find a large school of them. Gygarian trout are have a very distinct taste and are highly coveted. As you begin to haul in your nets the net is pulled hard. The crew looks over and runs, as you lean over the side of the ship a sharp tendril shoots through your stomach and a mighty kraken emerges. It appears he also has a taste for the trout and didn’t like you spoiling his catch. The crew cuts the nets and sails away as you hang from the tentacle. You try to fight it but your strength is drained from the injury. He tears you apart and the last thing you hear are your own bloody screams and the crunching of your bones."}

	// Ocean
	ocean := tbastory.MapCell{ID: 51, Title: "Ocean", Disc: "As you step on to the sandy shore the ocean laps at your feet. It reminds you of a romance you saw once, but all those people died on a boat anyways. Other than that nothing catches your interest."}

	// Cliff
	cliff := tbastory.MapCell{ID: 52, Title: "Cliff", Disc: "As you travel on the ocean, it leads to a large waterfall and a very large cliff. You doubt you can sail a boat up a cliff, so and as you progress further you see it now surrounds you to the north, east and west. You might as well go back until they invent some kind of flying ship."}

	// ***********
	// ** Conns **
	// ***********
	// Conns must be created after the intial creation of the Cells and then added to the Cells they belong in before adding the Cells to World
	// Each Conn must have a unique ID

	// Orc Camp (1)
	// SOUTH (if !wolves)
	conn1 := tbastory.MapConn{ID: 1, Name: "SOUTH", StartCellID: 1, EndCellID: 12, NeedMobKilled: true}
	conn1.RecMobIDs = append(conn1.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// SOUTH (if wolves)
	conn2 := tbastory.MapConn{ID: 2, Name: "SOUTH", StartCellID: 1, EndCellID: 11, NeedMobAlive: true}
	conn2.RecMobIDs = append(conn2.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// ATTACK (if orcs)
	conn3 := tbastory.MapConn{ID: 3, Name: "ATTACK", StartCellID: 1, EndCellID: 2, NeedMobAlive: true}
	conn3.RecMobIDs = append(conn3.RecMobIDs, orc1.ID, orc2.ID)
	// Add conns
	orcCamp.Conns = append(orcCamp.Conns, conn1, conn2, conn3)

	// Orc Camp (Empty) (2)
	// SOUTH (if !wolves)
	conn4 := tbastory.MapConn{ID: 4, Name: "SOUTH", StartCellID: 2, EndCellID: 12, NeedMobKilled: true}
	conn4.RecMobIDs = append(conn1.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// SOUTH (if wolves)
	conn5 := tbastory.MapConn{ID: 5, Name: "SOUTH", StartCellID: 2, EndCellID: 11, NeedMobAlive: true}
	conn5.RecMobIDs = append(conn2.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// Add conns
	orcCampEmpty.Conns = append(orcCampEmpty.Conns, conn4, conn5)

	// Castle Courtyard (3)
	// SOUTH
	conn6 := tbastory.MapConn{ID: 6, Name: "SOUTH", StartCellID: 3, EndCellID: 13}
	// EAST
	conn7 := tbastory.MapConn{ID: 7, Name: "EAST", StartCellID: 3, EndCellID: 7}
	// Add conns
	castleCourtyard.Conns = append(castleCourtyard.Conns, conn6, conn7)

	// Castle Courtyard (Seal) (4)
	// SOUTH (if sleep)
	conn8 := tbastory.MapConn{ID: 8, Name: "SOUTH", StartCellID: 4, EndCellID: 14, NeedMobAlive: true}
	conn8.RecMobIDs = append(conn6.RecMobIDs, sleep.ID)
	// SOUTH (if !sleep)
	conn9 := tbastory.MapConn{ID: 9, Name: "SOUTH", StartCellID: 4, EndCellID: 15, NeedMobKilled: true}
	conn9.RecMobIDs = append(conn7.RecMobIDs, sleep.ID)
	// EAST
	conn10 := tbastory.MapConn{ID: 10, Name: "EAST", StartCellID: 4, EndCellID: 7}
	// OPEN
	conn11 := tbastory.MapConn{ID: 11, Name: "OPEN", StartCellID: 4, EndCellID: 5, NeedItem: true}
	conn11.RecItemIDs = append(conn8.RecItemIDs, seal.ID)
	// Add conns
	castleCourtyardSeal.Conns = append(castleCourtyardSeal.Conns, conn8, conn9, conn10, conn11)

	// Castle (5)
	// ATTACK
	conn12 := tbastory.MapConn{ID: 12, Name: "ATTACK", StartCellID: 5, EndCellID: 6, NeedMobAlive: true}
	conn12.RecMobIDs = append(conn9.RecMobIDs, wizardAttack.ID)
	// Add conns
	castle.Conns = append(castle.Conns, conn12)

	// Land Slide (7)
	// EAST
	conn13 := tbastory.MapConn{ID: 13, Name: "EAST", StartCellID: 7, EndCellID: 8}
	// SOUTH (if demon)
	conn14 := tbastory.MapConn{ID: 14, Name: "SOUTH", StartCellID: 7, EndCellID: 16, NeedNotItem: true, NeedNPCNoTalk: true}
	conn14.RecItemIDs = append(conn14.RecItemIDs, potion.ID)
	conn14.RecNPCIDs = append(conn14.RecNPCIDs, demon.ID)
	// SOUTH (if demonAttack && !demonDeal)
	conn15 := tbastory.MapConn{ID: 15, Name: "SOUTH", StartCellID: 7, EndCellID: 17, NeedNotItem: true, NeedNPCNoTalk: true}
	conn15.RecItemIDs = append(conn15.RecItemIDs, potion.ID)
	conn15.RecNPCIDs = append(conn15.RecNPCIDs, demonDeal.ID)
	// SOUTH (if !demon)
	conn16 := tbastory.MapConn{ID: 16, Name: "SOUTH", StartCellID: 7, EndCellID: 18, NeedItem: true}
	conn16.RecItemIDs = append(conn16.RecItemIDs, potion.ID)
	// WEST (if !seal)
	conn17 := tbastory.MapConn{ID: 17, Name: "WEST", StartCellID: 7, EndCellID: 3, NeedNotItem: true}
	conn17.RecItemIDs = append(conn17.RecItemIDs, seal.ID)
	// WEST (if seal)
	conn18 := tbastory.MapConn{ID: 18, Name: "WEST", StartCellID: 7, EndCellID: 4, NeedItem: true}
	conn18.RecItemIDs = append(conn18.RecItemIDs, seal.ID)
	// Add conns
	landSlide.Conns = append(landSlide.Conns, conn13, conn14, conn15, conn16, conn17, conn18)

	// Dense Forest (8)
	// WEST
	conn19 := tbastory.MapConn{ID: 19, Name: "WEST", StartCellID: 8, EndCellID: 7}
	// Add Conns
	forest.Conns = append(forest.Conns, conn19)

	// Backwater Village (9)
	// SOUTH (if !cloak)
	conn20 := tbastory.MapConn{ID: 20, Name: "SOUTH", StartCellID: 9, EndCellID: 20, NeedNotItem: true}
	conn20.RecItemIDs = append(conn20.RecItemIDs, cloak.ID)
	// SOUTH (if cloak)
	conn21 := tbastory.MapConn{ID: 21, Name: "SOUTH", StartCellID: 9, EndCellID: 22, NeedItem: true}
	conn21.RecItemIDs = append(conn21.RecItemIDs, cloak.ID)
	// ATTACK (if cannibals)
	conn22 := tbastory.MapConn{ID: 22, Name: "ATTACK", StartCellID: 9, EndCellID: 10, NeedMobAlive: true}
	conn22.RecMobIDs = append(conn22.RecMobIDs, cannibal1.ID, cannibal2.ID, cannibal3.ID, cannibal4.ID, cannibal5.ID)
	// Add conns
	village.Conns = append(village.Conns, conn20, conn21, conn22)

	// Backwater Village (Empty) (10)
	// SOUTH (20) (if !cloak)
	conn23 := tbastory.MapConn{ID: 23, Name: "SOUTH", StartCellID: 10, EndCellID: 20, NeedNotItem: true}
	conn23.RecItemIDs = append(conn23.RecItemIDs, cloak.ID)
	// SOUTH (22) (if cloak)
	conn24 := tbastory.MapConn{ID: 24, Name: "SOUTH", StartCellID: 10, EndCellID: 22, NeedItem: true}
	conn24.RecItemIDs = append(conn24.RecItemIDs, cloak.ID)
	// Add conns
	villageEmpty.Conns = append(villageEmpty.Conns, conn23, conn24)

	// Wolves' Den (11)
	// NORTH (if orcs)
	conn25 := tbastory.MapConn{ID: 25, Name: "NORTH", StartCellID: 11, EndCellID: 1, NeedMobAlive: true}
	conn25.RecMobIDs = append(conn25.RecMobIDs, orc1.ID, orc2.ID)
	// NORTH (if !orcs)
	conn26 := tbastory.MapConn{ID: 26, Name: "NORTH", StartCellID: 11, EndCellID: 2, NeedMobKilled: true}
	conn26.RecMobIDs = append(conn26.RecMobIDs, orc1.ID, orc2.ID)
	// EAST (if !seal)
	conn27 := tbastory.MapConn{ID: 27, Name: "EAST", StartCellID: 11, EndCellID: 13, NeedNotItem: true}
	conn27.RecItemIDs = append(conn27.RecItemIDs, seal.ID)
	// EAST (if seal && sleap)
	conn28 := tbastory.MapConn{ID: 28, Name: "EAST", StartCellID: 11, EndCellID: 14, NeedItem: true, NeedMobAlive: true}
	conn28.RecItemIDs = append(conn28.RecItemIDs, seal.ID)
	conn28.RecMobIDs = append(conn28.RecMobIDs, sleep.ID)
	// EAST (if seal && !sleep)
	conn29 := tbastory.MapConn{ID: 29, Name: "EAST", StartCellID: 11, EndCellID: 15, NeedItem: true, NeedMobKilled: true}
	conn29.RecItemIDs = append(conn29.RecItemIDs, seal.ID)
	conn29.RecMobIDs = append(conn29.RecMobIDs, sleep.ID)
	// ATTACK (if wolves)
	conn30 := tbastory.MapConn{ID: 30, Name: "ATTACK", StartCellID: 11, EndCellID: 12, NeedMobAlive: true}
	conn30.RecMobIDs = append(conn30.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// Add conns
	den.Conns = append(den.Conns, conn25, conn26, conn27, conn28, conn29, conn30)

	// Wolves' Den (Empty) (12)
	// NORTH (if orcs)
	conn31 := tbastory.MapConn{ID: 31, Name: "NORTH", StartCellID: 12, EndCellID: 1, NeedMobAlive: true}
	conn31.RecMobIDs = append(conn31.RecMobIDs, orc1.ID, orc2.ID)
	// NORTH (if !orcs)
	conn32 := tbastory.MapConn{ID: 32, Name: "NORTH", StartCellID: 12, EndCellID: 2, NeedMobKilled: true}
	conn32.RecMobIDs = append(conn32.RecMobIDs, orc1.ID, orc2.ID)
	// EAST (if !seal)
	conn33 := tbastory.MapConn{ID: 33, Name: "EAST", StartCellID: 12, EndCellID: 13, NeedNotItem: true}
	conn33.RecItemIDs = append(conn33.RecItemIDs, seal.ID)
	// EAST (if seal && sleap)
	conn34 := tbastory.MapConn{ID: 34, Name: "EAST", StartCellID: 12, EndCellID: 14, NeedItem: true, NeedMobAlive: true}
	conn34.RecItemIDs = append(conn34.RecItemIDs, seal.ID)
	conn34.RecMobIDs = append(conn34.RecMobIDs, sleep.ID)
	// EAST (if seal && !sleep)
	conn35 := tbastory.MapConn{ID: 35, Name: "EAST", StartCellID: 12, EndCellID: 15, NeedItem: true, NeedMobKilled: true}
	conn35.RecItemIDs = append(conn35.RecItemIDs, seal.ID)
	conn35.RecMobIDs = append(conn35.RecMobIDs, sleep.ID)
	// Add conns
	denEmpty.Conns = append(denEmpty.Conns, conn31, conn32, conn33, conn34, conn35)

	// Spring Meadow (13)
	// NORTH
	conn41 := tbastory.MapConn{ID: 41, Name: "NORTH", StartCellID: 13, EndCellID: 3}
	// EAST (if demon)
	conn36 := tbastory.MapConn{ID: 36, Name: "EAST", StartCellID: 13, EndCellID: 16, NeedNotItem: true, NeedNPCNoTalk: true}
	conn36.RecItemIDs = append(conn36.RecItemIDs, potion.ID)
	conn36.RecNPCIDs = append(conn36.RecNPCIDs, demon.ID)
	// EAST (if demonAttack && !demonDeal)
	conn37 := tbastory.MapConn{ID: 37, Name: "EAST", StartCellID: 13, EndCellID: 17, NeedNotItem: true, NeedNPCNoTalk: true}
	conn37.RecItemIDs = append(conn37.RecItemIDs, potion.ID)
	conn37.RecNPCIDs = append(conn37.RecNPCIDs, demonDeal.ID)
	// EAST (if !demon)
	conn38 := tbastory.MapConn{ID: 38, Name: "EAST", StartCellID: 13, EndCellID: 18, NeedItem: true}
	conn38.RecItemIDs = append(conn38.RecItemIDs, potion.ID)
	// WEST (if !wolves)
	conn39 := tbastory.MapConn{ID: 39, Name: "WEST", StartCellID: 13, EndCellID: 12, NeedMobKilled: true}
	conn39.RecMobIDs = append(conn39.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// WEST (if wolves)
	conn40 := tbastory.MapConn{ID: 40, Name: "WEST", StartCellID: 13, EndCellID: 11, NeedMobAlive: true}
	conn40.RecMobIDs = append(conn40.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// Add conns
	meadow.Conns = append(meadow.Conns, conn41, conn36, conn37, conn38, conn39, conn40)

	// Spring Meadow (Seal) (14)
	// ATTACK (if sleep)
	conn42 := tbastory.MapConn{ID: 42, Name: "ATTACK", StartCellID: 14, EndCellID: 15, NeedMobAlive: true}
	conn42.RecMobIDs = append(conn42.RecMobIDs, sleep.ID)
	// Add conns
	meadowSeal.Conns = append(meadowSeal.Conns, conn42)

	// Spring Meadow (Empty) (15)
	// NORTH
	conn43 := tbastory.MapConn{ID: 43, Name: "NORTH", StartCellID: 15, EndCellID: 4}
	// EAST (if demon)
	conn44 := tbastory.MapConn{ID: 44, Name: "EAST", StartCellID: 15, EndCellID: 16, NeedNotItem: true, NeedNPCNoTalk: true}
	conn44.RecItemIDs = append(conn44.RecItemIDs, potion.ID)
	conn44.RecNPCIDs = append(conn44.RecNPCIDs, demon.ID)
	// EAST (if demonAttack && !demonDeal)
	conn45 := tbastory.MapConn{ID: 45, Name: "EAST", StartCellID: 15, EndCellID: 17, NeedNotItem: true, NeedNPCNoTalk: true}
	conn45.RecItemIDs = append(conn45.RecItemIDs, potion.ID)
	conn45.RecNPCIDs = append(conn45.RecNPCIDs, demonDeal.ID)
	// EAST (if !demon)
	conn46 := tbastory.MapConn{ID: 46, Name: "EAST", StartCellID: 15, EndCellID: 18, NeedItem: true}
	conn46.RecItemIDs = append(conn46.RecItemIDs, potion.ID)
	// WEST (if !wolves)
	conn47 := tbastory.MapConn{ID: 47, Name: "WEST", StartCellID: 15, EndCellID: 12, NeedMobKilled: true}
	conn47.RecMobIDs = append(conn47.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// WEST (if wolves)
	conn48 := tbastory.MapConn{ID: 48, Name: "WEST", StartCellID: 15, EndCellID: 11, NeedMobAlive: true}
	conn48.RecMobIDs = append(conn48.RecMobIDs, wolf1.ID, wolf2.ID, wolf3.ID, wolf4.ID)
	// Add conns
	meadowEmpty.Conns = append(meadowEmpty.Conns, conn43, conn44, conn45, conn46, conn47, conn48)

	// Crossroads (16)
	// NORTH
	conn49 := tbastory.MapConn{ID: 49, Name: "NORTH", StartCellID: 16, EndCellID: 7}
	// EAST
	conn50 := tbastory.MapConn{ID: 50, Name: "EAST", StartCellID: 16, EndCellID: 19}
	// SOUTH
	conn51 := tbastory.MapConn{ID: 51, Name: "SOUTH", StartCellID: 16, EndCellID: 29}
	// WEST (if !seal)
	conn52 := tbastory.MapConn{ID: 52, Name: "WEST", StartCellID: 16, EndCellID: 13, NeedNotItem: true}
	conn52.RecItemIDs = append(conn52.RecItemIDs, seal.ID)
	// WEST (if seal && sleap)
	conn53 := tbastory.MapConn{ID: 53, Name: "WEST", StartCellID: 16, EndCellID: 14, NeedItem: true, NeedMobAlive: true}
	conn53.RecItemIDs = append(conn53.RecItemIDs, seal.ID)
	conn53.RecMobIDs = append(conn53.RecMobIDs, sleep.ID)
	// WEST (if seal && !sleep)
	conn54 := tbastory.MapConn{ID: 54, Name: "WEST", StartCellID: 15, EndCellID: 15, NeedItem: true, NeedMobKilled: true}
	conn54.RecItemIDs = append(conn54.RecItemIDs, seal.ID)
	conn54.RecMobIDs = append(conn54.RecMobIDs, sleep.ID)
	// ATTACK (if demonAttack)
	conn55 := tbastory.MapConn{ID: 55, Name: "ATTACK", StartCellID: 16, EndCellID: 18, NeedMobAlive: true}
	conn55.RecMobIDs = append(conn55.RecMobIDs, demonAttack.ID)
	// TALK (if demon)
	conn56 := tbastory.MapConn{ID: 56, Name: "TALK", StartCellID: 16, EndCellID: 17, NeedNPCNoTalk: true}
	conn56.RecNPCIDs = append(conn56.RecNPCIDs, demon.ID)
	// Add conns
	crossroads.Conns = append(crossroads.Conns, conn49, conn50, conn51, conn52, conn53, conn54, conn55, conn56)

	// Crossroads (Demon) (17)
	// NORTH
	conn57 := tbastory.MapConn{ID: 57, Name: "NORTH", StartCellID: 17, EndCellID: 7}
	// EAST
	conn58 := tbastory.MapConn{ID: 58, Name: "EAST", StartCellID: 17, EndCellID: 19}
	// SOUTH
	conn59 := tbastory.MapConn{ID: 59, Name: "SOUTH", StartCellID: 17, EndCellID: 29}
	// WEST (if !seal)
	conn60 := tbastory.MapConn{ID: 60, Name: "WEST", StartCellID: 17, EndCellID: 13, NeedNotItem: true}
	conn60.RecItemIDs = append(conn60.RecItemIDs, seal.ID)
	// WEST (if seal && sleap)
	conn61 := tbastory.MapConn{ID: 61, Name: "WEST", StartCellID: 17, EndCellID: 14, NeedItem: true, NeedMobAlive: true}
	conn61.RecItemIDs = append(conn61.RecItemIDs, seal.ID)
	conn61.RecMobIDs = append(conn61.RecMobIDs, sleep.ID)
	// WEST (if seal && !sleep)
	conn62 := tbastory.MapConn{ID: 62, Name: "WEST", StartCellID: 17, EndCellID: 15, NeedItem: true, NeedMobKilled: true}
	conn62.RecItemIDs = append(conn62.RecItemIDs, seal.ID)
	conn62.RecMobIDs = append(conn62.RecMobIDs, sleep.ID)
	// ATTACK (if demonAttack)
	conn63 := tbastory.MapConn{ID: 63, Name: "ATTACK", StartCellID: 17, EndCellID: 18, NeedMobAlive: true}
	conn63.RecMobIDs = append(conn63.RecMobIDs, demonAttack.ID)
	// TALK (if demon)
	conn64 := tbastory.MapConn{ID: 64, Name: "TALK", StartCellID: 17, EndCellID: 18, NeedNPCNoTalk: true}
	conn64.RecNPCIDs = append(conn64.RecNPCIDs, demonDeal.ID)
	// Add conns
	crossroadsDeal.Conns = append(crossroadsDeal.Conns, conn57, conn58, conn59, conn60, conn61, conn62, conn63, conn64)

	// Crossroads (Empty) (18)
	// NORTH
	conn65 := tbastory.MapConn{ID: 65, Name: "NORTH", StartCellID: 18, EndCellID: 7}
	// EAST
	conn66 := tbastory.MapConn{ID: 66, Name: "EAST", StartCellID: 18, EndCellID: 19}
	// SOUTH
	conn67 := tbastory.MapConn{ID: 67, Name: "SOUTH", StartCellID: 18, EndCellID: 29}
	// WEST (if !seal)
	conn68 := tbastory.MapConn{ID: 68, Name: "WEST", StartCellID: 18, EndCellID: 13, NeedNotItem: true}
	conn68.RecItemIDs = append(conn60.RecItemIDs, seal.ID)
	// WEST (if seal && sleap)
	conn69 := tbastory.MapConn{ID: 69, Name: "WEST", StartCellID: 18, EndCellID: 14, NeedItem: true, NeedMobAlive: true}
	conn69.RecItemIDs = append(conn69.RecItemIDs, seal.ID)
	conn69.RecMobIDs = append(conn69.RecMobIDs, sleep.ID)
	// WEST (if seal && !sleep)
	conn70 := tbastory.MapConn{ID: 70, Name: "WEST", StartCellID: 18, EndCellID: 15, NeedItem: true, NeedMobKilled: true}
	conn70.RecItemIDs = append(conn70.RecItemIDs, seal.ID)
	conn70.RecMobIDs = append(conn70.RecMobIDs, sleep.ID)
	// Add conns
	crossroadsEmpty.Conns = append(crossroadsEmpty.Conns, conn65, conn66, conn67, conn68, conn69, conn70)

	// Bridge (19)
	// EAST (if !cloak)
	conn71 := tbastory.MapConn{ID: 71, Name: "EAST", StartCellID: 19, EndCellID: 20, NeedNotItem: true}
	conn71.RecItemIDs = append(conn71.RecItemIDs, cloak.ID)
	// EAST (if cloak)
	conn72 := tbastory.MapConn{ID: 72, Name: "EAST", StartCellID: 19, EndCellID: 22, NeedItem: true}
	conn72.RecItemIDs = append(conn72.RecItemIDs, cloak.ID)
	// WEST (if demon)
	conn73 := tbastory.MapConn{ID: 73, Name: "WEST", StartCellID: 19, EndCellID: 16, NeedNotItem: true, NeedNPCNoTalk: true}
	conn73.RecItemIDs = append(conn73.RecItemIDs, potion.ID)
	conn73.RecNPCIDs = append(conn73.RecNPCIDs, demon.ID)
	// WEST (if demonAttack && !demonDeal)
	conn74 := tbastory.MapConn{ID: 74, Name: "WEST", StartCellID: 19, EndCellID: 17, NeedNotItem: true, NeedNPCNoTalk: true}
	conn74.RecItemIDs = append(conn74.RecItemIDs, potion.ID)
	conn74.RecNPCIDs = append(conn74.RecNPCIDs, demonDeal.ID)
	// WEST (if !demon)
	conn75 := tbastory.MapConn{ID: 75, Name: "WEST", StartCellID: 19, EndCellID: 18, NeedItem: true}
	conn75.RecItemIDs = append(conn75.RecItemIDs, potion.ID)
	// Add conns
	bridge.Conns = append(bridge.Conns, conn71, conn72, conn73, conn74, conn75)

	// Boneyard (20)
	// NORTH (if cannibals)
	conn76 := tbastory.MapConn{ID: 76, Name: "NORTH", StartCellID: 20, EndCellID: 9, NeedMobAlive: true}
	conn76.RecMobIDs = append(conn76.RecMobIDs, cannibal1.ID, cannibal2.ID, cannibal3.ID, cannibal4.ID, cannibal5.ID)
	// NORTH (if !cannibals)
	conn77 := tbastory.MapConn{ID: 77, Name: "NORTH", StartCellID: 20, EndCellID: 10, NeedMobKilled: true}
	conn77.RecMobIDs = append(conn77.RecMobIDs, cannibal1.ID, cannibal2.ID, cannibal3.ID, cannibal4.ID, cannibal5.ID)
	// WEST
	conn78 := tbastory.MapConn{ID: 78, Name: "WEST", StartCellID: 20, EndCellID: 19}
	// JOIN
	conn79 := tbastory.MapConn{ID: 79, Name: "JOIN", StartCellID: 20, EndCellID: 23}
	// GRAB (if !cloak)
	conn80 := tbastory.MapConn{ID: 80, Name: "GRAB", StartCellID: 20, EndCellID: 21, NeedNotItem: true}
	conn80.RecItemIDs = append(conn80.RecItemIDs, cloak.ID)
	// Add conns
	boneyard.Conns = append(boneyard.Conns, conn76, conn77, conn78, conn79, conn80)

	// Boneyard (Grab) (21)
	// ATTACK
	conn81 := tbastory.MapConn{ID: 81, Name: "ATTACK", StartCellID: 21, EndCellID: 22, NeedMobAlive: true}
	conn81.RecMobIDs = append(conn81.RecMobIDs, skele1.ID, skele2.ID)
	// Add conns
	boneyardGrab.Conns = append(boneyardGrab.Conns, conn81)

	// Boneyard (Empty) (22)
	// NORTH (if cannibals)
	conn82 := tbastory.MapConn{ID: 82, Name: "NORTH", StartCellID: 22, EndCellID: 9, NeedMobAlive: true}
	conn82.RecMobIDs = append(conn82.RecMobIDs, cannibal1.ID, cannibal2.ID, cannibal3.ID, cannibal4.ID, cannibal5.ID)
	// NORTH (if !cannibals)
	conn83 := tbastory.MapConn{ID: 83, Name: "NORTH", StartCellID: 22, EndCellID: 10, NeedMobKilled: true}
	conn83.RecMobIDs = append(conn77.RecMobIDs, cannibal1.ID, cannibal2.ID, cannibal3.ID, cannibal4.ID, cannibal5.ID)
	// WEST
	conn84 := tbastory.MapConn{ID: 84, Name: "WEST", StartCellID: 22, EndCellID: 19}
	// JOIN
	conn85 := tbastory.MapConn{ID: 85, Name: "JOIN", StartCellID: 22, EndCellID: 23}
	// Add conns
	boneyardEmpty.Conns = append(boneyardEmpty.Conns, conn82, conn83, conn84, conn85)

	// Fairy Glade (24)
	// EAST
	conn86 := tbastory.MapConn{ID: 86, Name: "EAST", StartCellID: 24, EndCellID: 27}
	// TALK
	conn87 := tbastory.MapConn{ID: 86, Name: "TALK", StartCellID: 24, EndCellID: 24, NeedNPCNoTalk: true}
	conn87.RecNPCIDs = append(conn87.RecNPCIDs, fairy.ID)
	// Add conns
	glade.Conns = append(glade.Conns, conn86, conn87)

	// Fairy Glade (Quest) (25)
	// EAST
	conn88 := tbastory.MapConn{ID: 88, Name: "EAST", StartCellID: 25, EndCellID: 27}
	// TALK (if !potion)
	conn89 := tbastory.MapConn{ID: 89, Name: "TALK", StartCellID: 25, EndCellID: 25, NeedNotItem: true}
	conn89.RecItemIDs = append(conn89.RecItemIDs, potion.ID)
	// Add conns
	gladeQuest.Conns = append(gladeQuest.Conns, conn88, conn89)

	// Fairy Glade (Potion) (26)
	// EAST
	conn90 := tbastory.MapConn{ID: 1.0, Name: "EAST", StartCellID: 26, EndCellID: 27}
	// TALK (if potion)
	conn91 := tbastory.MapConn{ID: 1.0, Name: "TALK", StartCellID: 26, EndCellID: 24, NeedItem: true}
	conn91.RecItemIDs = append(conn91.RecItemIDs, potion.ID)
	// Add conns
	gladePotion.Conns = append(gladePotion.Conns, conn90, conn91)

	// Crystal Lake (27)
	// EAST
	conn92 := tbastory.MapConn{ID: 92, Name: "EAST", StartCellID: 27, EndCellID: 29}
	// WEST (if fairyPotion && potion)
	conn93 := tbastory.MapConn{ID: 93, Name: "WEST", StartCellID: 27, EndCellID: 24, NeedItem: true, NeedNPCTalk: true}
	conn93.RecItemIDs = append(conn93.RecItemIDs, potion.ID)
	conn93.RecNPCIDs = append(conn93.RecNPCIDs, fairyPotion.ID)
	// WEST (if !potion && !fairyBefore)
	conn94 := tbastory.MapConn{ID: 94, Name: "WEST", StartCellID: 27, EndCellID: 25, NeedNotItem: true}
	conn94.RecItemIDs = append(conn94.RecItemIDs, potion.ID)
	// WEST (if potion)
	conn95 := tbastory.MapConn{ID: 95, Name: "WEST", StartCellID: 27, EndCellID: 26, NeedItem: true, NeedNPCNoTalk: true}
	conn95.RecItemIDs = append(conn95.RecItemIDs, potion.ID)
	conn95.RecNPCIDs = append(conn95.RecNPCIDs, fairyPotion.ID)
	// Add conns
	lake.Conns = append(lake.Conns, conn92, conn93, conn94, conn95)

	// Crystal Lake (Quest) (28)
	// EAST (29)
	conn96 := tbastory.MapConn{ID: 96, Name: "EAST", StartCellID: 28, EndCellID: 25}
	// TALK (28) (if !coin)
	conn97 := tbastory.MapConn{ID: 97, Name: "TALK", StartCellID: 28, EndCellID: 28, NeedNotItem: true}
	conn97.RecItemIDs = append(conn97.RecItemIDs, coin.ID)
	// Add conns
	lakeQuest.Conns = append(lakeQuest.Conns, conn96, conn97)

	// Caravan (29)
	// NORTH (if demon)
	conn98 := tbastory.MapConn{ID: 98, Name: "NORTH", StartCellID: 29, EndCellID: 16, NeedNotItem: true, NeedNPCNoTalk: true}
	conn98.RecItemIDs = append(conn98.RecItemIDs, potion.ID)
	conn98.RecNPCIDs = append(conn98.RecNPCIDs, demon.ID)
	// NORTH (if demonAttack && !demonDeal)
	conn99 := tbastory.MapConn{ID: 99, Name: "NORTH", StartCellID: 29, EndCellID: 17, NeedNotItem: true, NeedNPCNoTalk: true}
	conn99.RecItemIDs = append(conn99.RecItemIDs, potion.ID)
	conn99.RecNPCIDs = append(conn99.RecNPCIDs, demonDeal.ID)
	// NORTH (if !demon)
	conn100 := tbastory.MapConn{ID: 100, Name: "NORTH", StartCellID: 29, EndCellID: 18, NeedItem: true}
	conn100.RecItemIDs = append(conn100.RecItemIDs, potion.ID)
	// EAST
	conn101 := tbastory.MapConn{ID: 101, Name: "EAST", StartCellID: 29, EndCellID: 30}
	// SOUTH (if oldMan && lady)
	conn102 := tbastory.MapConn{ID: 102, Name: "SOUTH", StartCellID: 29, EndCellID: 38, NeedNPCTalk: true}
	conn102.RecNPCIDs = append(conn102.RecNPCIDs, oldMan.ID, lady.ID)
	// SOUTH (if !armor)
	conn103 := tbastory.MapConn{ID: 103, Name: "SOUTH", StartCellID: 29, EndCellID: 39, NeedNotItem: true}
	conn103.RecItemIDs = append(conn103.RecItemIDs, armor.ID)
	// SOUTH (if !sword && !cloak)
	conn104 := tbastory.MapConn{ID: 104, Name: "SOUTH", StartCellID: 29, EndCellID: 40, NeedNotItem: true}
	conn104.RecItemIDs = append(conn104.RecItemIDs, sword.ID, cloak.ID)
	// WEST (if coin)
	conn105 := tbastory.MapConn{ID: 105, Name: "WEST", StartCellID: 29, EndCellID: 27, NeedItem: true}
	conn105.RecItemIDs = append(conn105.RecItemIDs, coin.ID)
	// WEST (if !coin)
	conn106 := tbastory.MapConn{ID: 106, Name: "WEST", StartCellID: 29, EndCellID: 28, NeedNotItem: true}
	conn106.RecItemIDs = append(conn106.RecItemIDs, coin.ID)
	// Add conns
	caravan.Conns = append(caravan.Conns, conn98, conn99, conn100, conn101, conn102, conn103, conn104, conn105, conn106)

	// Plains (30)
	// EAST (if sandworm)
	conn107 := tbastory.MapConn{ID: 107, Name: "EAST", StartCellID: 30, EndCellID: 31, NeedMobAlive: true}
	conn107.RecMobIDs = append(conn107.RecMobIDs, sandworm.ID)
	// EAST (if !sandworm)
	conn108 := tbastory.MapConn{ID: 108, Name: "EAST", StartCellID: 30, EndCellID: 32, NeedMobKilled: true}
	conn108.RecMobIDs = append(conn108.RecMobIDs, sandworm.ID)
	// SOUTH
	conn109 := tbastory.MapConn{ID: 109, Name: "SOUTH", StartCellID: 30, EndCellID: 41}
	// WEST
	conn110 := tbastory.MapConn{ID: 110, Name: "WEST", StartCellID: 30, EndCellID: 29}
	// Add conns
	plains.Conns = append(plains.Conns, conn107, conn108, conn109, conn110)

	// Desert (31)
	// WEST
	conn111 := tbastory.MapConn{ID: 111, Name: "WEST", StartCellID: 31, EndCellID: 30}
	// ATTACK (if sandworm)
	conn112 := tbastory.MapConn{ID: 112, Name: "ATTACK", StartCellID: 31, EndCellID: 32, NeedMobAlive: true}
	conn112.RecMobIDs = append(conn112.RecMobIDs, sandworm.ID)
	// Add conns
	desert.Conns = append(desert.Conns, conn111, conn112)

	// Desert (Empty) (32)
	// WEST
	conn113 := tbastory.MapConn{ID: 113, Name: "WEST", StartCellID: 32, EndCellID: 30}
	// Add conns
	desertEmpty.Conns = append(desertEmpty.Conns, conn113)

	// Grave Keeper's Hovel (33)
	// EAST (if zombie)
	conn114 := tbastory.MapConn{ID: 114, Name: "EAST", StartCellID: 33, EndCellID: 35, NeedMobAlive: true}
	conn114.RecMobIDs = append(conn114.RecMobIDs, zombie.ID)
	// EAST (if !zombie)
	conn115 := tbastory.MapConn{ID: 115, Name: "EAST", StartCellID: 33, EndCellID: 37, NeedMobKilled: true}
	conn115.RecMobIDs = append(conn115.RecMobIDs, zombie.ID)
	// SOUTH
	conn116 := tbastory.MapConn{ID: 116, Name: "SOUTH", StartCellID: 33, EndCellID: 44}
	// TAKE
	conn117 := tbastory.MapConn{ID: 117, Name: "SOUTH", StartCellID: 33, EndCellID: 34}
	// Add conns
	hovel.Conns = append(hovel.Conns, conn114, conn115, conn116, conn117)

	// Town Graveyard (35)
	// EAST (if oldMan && lady)
	conn118 := tbastory.MapConn{ID: 118, Name: "EAST", StartCellID: 35, EndCellID: 38, NeedNPCTalk: true}
	conn118.RecNPCIDs = append(conn118.RecNPCIDs, oldMan.ID, lady.ID)
	// EAST (if !armor)
	conn119 := tbastory.MapConn{ID: 119, Name: "EAST", StartCellID: 35, EndCellID: 39, NeedNotItem: true}
	conn119.RecItemIDs = append(conn119.RecItemIDs, armor.ID)
	// EAST (if !sword && !cloak)
	conn120 := tbastory.MapConn{ID: 120, Name: "EAST", StartCellID: 35, EndCellID: 40, NeedNotItem: true}
	conn120.RecItemIDs = append(conn120.RecItemIDs, sword.ID, cloak.ID)
	// WEST
	conn121 := tbastory.MapConn{ID: 121, Name: "WEST", StartCellID: 35, EndCellID: 33}
	// INSPECT
	conn122 := tbastory.MapConn{ID: 122, Name: "INSPECT", StartCellID: 35, EndCellID: 36, NeedMobAlive: true}
	conn122.RecMobIDs = append(conn122.RecMobIDs, zombie.ID)
	// Add conns
	graveyard.Conns = append(graveyard.Conns, conn118, conn119, conn120, conn121, conn122)

	// Town Graveyard (Inspect) (36)
	// ATTACK
	conn123 := tbastory.MapConn{ID: 123, Name: "ATTACK", StartCellID: 36, EndCellID: 37, NeedMobAlive: true}
	conn123.RecMobIDs = append(conn123.RecMobIDs, zombie.ID)
	// Add conns
	graveyardInspect.Conns = append(graveyardInspect.Conns, conn123)

	// Town Graveyard (Empty) (37)
	// EAST (if oldMan && lady)
	conn124 := tbastory.MapConn{ID: 124, Name: "EAST", StartCellID: 37, EndCellID: 38, NeedNPCTalk: true}
	conn124.RecNPCIDs = append(conn124.RecNPCIDs, oldMan.ID, lady.ID)
	// EAST (if !armor)
	conn125 := tbastory.MapConn{ID: 125, Name: "EAST", StartCellID: 37, EndCellID: 39, NeedNotItem: true}
	conn125.RecItemIDs = append(conn125.RecItemIDs, armor.ID)
	// EAST (if !sword && !cloak)
	conn126 := tbastory.MapConn{ID: 126, Name: "EAST", StartCellID: 37, EndCellID: 40, NeedNotItem: true}
	conn126.RecItemIDs = append(conn126.RecItemIDs, sword.ID, cloak.ID)
	// WEST
	conn127 := tbastory.MapConn{ID: 127, Name: "WEST", StartCellID: 37, EndCellID: 33}
	// Add conns
	graveyardEmpty.Conns = append(graveyardEmpty.Conns, conn124, conn125, conn126, conn127)

	// Town (38)
	// NORTH
	conn128 := tbastory.MapConn{ID: 128, Name: "NORTH", StartCellID: 38, EndCellID: 29}
	// EAST
	conn129 := tbastory.MapConn{ID: 129, Name: "EAST", StartCellID: 38, EndCellID: 41}
	// SOUTH
	conn130 := tbastory.MapConn{ID: 130, Name: "SOUTH", StartCellID: 38, EndCellID: 49}
	// WEST (if zombie)
	conn131 := tbastory.MapConn{ID: 131, Name: "WEST", StartCellID: 38, EndCellID: 35, NeedMobAlive: true}
	conn131.RecMobIDs = append(conn131.RecMobIDs, zombie.ID)
	// WEST (if !zombie)
	conn132 := tbastory.MapConn{ID: 132, Name: "WEST", StartCellID: 38, EndCellID: 37, NeedMobKilled: true}
	conn132.RecMobIDs = append(conn132.RecMobIDs, zombie.ID)
	// Add conns
	town.Conns = append(town.Conns, conn128, conn129, conn130, conn131, conn132)

	// Town (Old Man) (39)
	// NORTH
	conn133 := tbastory.MapConn{ID: 133, Name: "NORTH", StartCellID: 39, EndCellID: 29}
	// EAST
	conn134 := tbastory.MapConn{ID: 134, Name: "EAST", StartCellID: 39, EndCellID: 41}
	// SOUTH
	conn135 := tbastory.MapConn{ID: 135, Name: "SOUTH", StartCellID: 39, EndCellID: 49}
	// WEST (if zombie)
	conn136 := tbastory.MapConn{ID: 136, Name: "WEST", StartCellID: 39, EndCellID: 35, NeedMobAlive: true}
	conn136.RecMobIDs = append(conn136.RecMobIDs, zombie.ID)
	// WEST (if !zombie)
	conn137 := tbastory.MapConn{ID: 137, Name: "WEST", StartCellID: 39, EndCellID: 37, NeedMobKilled: true}
	conn137.RecMobIDs = append(conn137.RecMobIDs, zombie.ID)
	// TALK (if !oldMan)
	conn138 := tbastory.MapConn{ID: 138, Name: "TALK", StartCellID: 39, EndCellID: 38, NeedNPCNoTalk: true}
	conn138.RecNPCIDs = append(conn138.RecNPCIDs, oldMan.ID)
	// Add conns
	townMan.Conns = append(townMan.Conns, conn133, conn134, conn135, conn136, conn137, conn138)

	// Town (Lady in Rags) (40)
	// NORTH
	conn139 := tbastory.MapConn{ID: 139, Name: "NORTH", StartCellID: 40, EndCellID: 29}
	// EAST
	conn140 := tbastory.MapConn{ID: 140, Name: "EAST", StartCellID: 40, EndCellID: 41}
	// SOUTH
	conn141 := tbastory.MapConn{ID: 141, Name: "SOUTH", StartCellID: 40, EndCellID: 49}
	// WEST (if zombie)
	conn142 := tbastory.MapConn{ID: 142, Name: "WEST", StartCellID: 40, EndCellID: 35, NeedMobAlive: true}
	conn142.RecMobIDs = append(conn142.RecMobIDs, zombie.ID)
	// WEST (if !zombie)
	conn143 := tbastory.MapConn{ID: 143, Name: "WEST", StartCellID: 40, EndCellID: 37, NeedMobKilled: true}
	conn143.RecMobIDs = append(conn143.RecMobIDs, zombie.ID)
	// TALK (if !oldMan)
	conn144 := tbastory.MapConn{ID: 144, Name: "TALK", StartCellID: 40, EndCellID: 38, NeedNPCNoTalk: true}
	conn144.RecNPCIDs = append(conn144.RecNPCIDs, lady.ID)
	// Add conns
	townLady.Conns = append(townLady.Conns, conn139, conn140, conn141, conn142, conn143, conn144)

	// Farm (41)
	// NORTH
	conn145 := tbastory.MapConn{ID: 145, Name: "NORTH", StartCellID: 41, EndCellID: 30}
	// EAST (if spider)
	conn146 := tbastory.MapConn{ID: 146, Name: "EAST", StartCellID: 41, EndCellID: 42, NeedMobAlive: true}
	conn146.RecMobIDs = append(conn146.RecMobIDs, spider.ID)
	// EAST (if !spider)
	conn147 := tbastory.MapConn{ID: 147, Name: "EAST", StartCellID: 41, EndCellID: 43, NeedMobKilled: true}
	conn147.RecMobIDs = append(conn147.RecMobIDs, spider.ID)
	// SOUTH
	conn148 := tbastory.MapConn{ID: 148, Name: "SOUTH", StartCellID: 41, EndCellID: 51}
	// WEST (if oldMan && lady)
	conn149 := tbastory.MapConn{ID: 149, Name: "WEST", StartCellID: 41, EndCellID: 38, NeedNPCTalk: true}
	conn149.RecNPCIDs = append(conn149.RecNPCIDs, oldMan.ID, lady.ID)
	// WEST (if !armor)
	conn150 := tbastory.MapConn{ID: 150, Name: "WEST", StartCellID: 41, EndCellID: 39, NeedNotItem: true}
	conn150.RecItemIDs = append(conn150.RecItemIDs, armor.ID)
	// WEST (if !sword && !cloak)
	conn151 := tbastory.MapConn{ID: 151, Name: "WEST", StartCellID: 41, EndCellID: 40, NeedNotItem: true}
	conn151.RecItemIDs = append(conn151.RecItemIDs, sword.ID, cloak.ID)
	// Add conns
	farm.Conns = append(farm.Conns, conn145, conn146, conn147, conn148, conn149, conn150, conn151)

	// Spider Forest (42)
	// WEST
	conn152 := tbastory.MapConn{ID: 152, Name: "WEST", StartCellID: 42, EndCellID: 41}
	// ATTACK (if spider)
	conn153 := tbastory.MapConn{ID: 153, Name: "ATTACK", StartCellID: 42, EndCellID: 43, NeedMobAlive: true}
	conn153.RecMobIDs = append(conn153.RecMobIDs, spider.ID)
	// Add conns
	spiderForest.Conns = append(spiderForest.Conns, conn152, conn153)

	// Spider Forest (Empty) (43)
	// WEST
	conn154 := tbastory.MapConn{ID: 154, Name: "WEST", StartCellID: 43, EndCellID: 41}
	// Add conns
	spiderForestEmpty.Conns = append(spiderForestEmpty.Conns, conn154)

	// Mountain (44)
	// NORTH
	conn155 := tbastory.MapConn{ID: 155, Name: "NORTH", StartCellID: 44, EndCellID: 33}
	// Add conns
	mountain.Conns = append(mountain.Conns, conn155)

	// Dragon's Lair (45)
	// EAST
	conn156 := tbastory.MapConn{ID: 156, Name: "EAST", StartCellID: 45, EndCellID: 49}
	// TALK
	conn157 := tbastory.MapConn{ID: 157, Name: "TALK", StartCellID: 45, EndCellID: 46, NeedNotItem: true}
	conn157.RecItemIDs = append(conn157.RecItemIDs, sword.ID)
	// SNEAK (if cloak)
	conn158 := tbastory.MapConn{ID: 158, Name: "SNEAK", StartCellID: 45, EndCellID: 47, NeedItem: true}
	conn158.RecItemIDs = append(conn157.RecItemIDs, cloak.ID)
	// ATTACK
	conn159 := tbastory.MapConn{ID: 159, Name: "ATTACK", StartCellID: 45, EndCellID: 48, NeedMobAlive: true}
	conn159.RecMobIDs = append(conn159.RecMobIDs, dragonAttack.ID)
	// Add conns
	dragonLair.Conns = append(dragonLair.Conns, conn156, conn157, conn158, conn159)

	// Dragon's Lair (Riddle) (46)
	// A
	conn160 := tbastory.MapConn{ID: 160, Name: "ANSWER A", StartCellID: 46, EndCellID: 0}
	// B
	conn161 := tbastory.MapConn{ID: 161, Name: "ANSWER B", StartCellID: 46, EndCellID: 48}
	// C
	conn162 := tbastory.MapConn{ID: 162, Name: "ANSWER C", StartCellID: 46, EndCellID: 0}
	// D
	conn163 := tbastory.MapConn{ID: 163, Name: "ANSWER D", StartCellID: 46, EndCellID: 0}
	// Add conns
	dragonLairRiddle.Conns = append(dragonLairRiddle.Conns, conn160, conn161, conn162, conn163)

	// Dragon's Lair (Sneak) (47)
	// EAST (27)
	conn164 := tbastory.MapConn{ID: 164, Name: "EAST", StartCellID: 47, EndCellID: 49}
	// Add conns
	dragonLairSneak.Conns = append(dragonLairSneak.Conns, conn164)

	// Dragon's Lair (Empty) (48)
	// EAST (27)
	conn165 := tbastory.MapConn{ID: 165, Name: "EAST", StartCellID: 48, EndCellID: 49}
	// Add conns
	dragonLairEmpty.Conns = append(dragonLairEmpty.Conns, conn165)

	// Docks (49)
	// NORTH (if oldMan && lady)
	conn166 := tbastory.MapConn{ID: 166, Name: "NORTH", StartCellID: 49, EndCellID: 38, NeedNPCTalk: true}
	conn166.RecNPCIDs = append(conn166.RecNPCIDs, oldMan.ID, lady.ID)
	// NORTH (if !armor)
	conn167 := tbastory.MapConn{ID: 167, Name: "NORTH", StartCellID: 49, EndCellID: 39, NeedNotItem: true}
	conn167.RecItemIDs = append(conn167.RecItemIDs, armor.ID)
	// NORTH (if !sword && !cloak)
	conn168 := tbastory.MapConn{ID: 168, Name: "NORTH", StartCellID: 49, EndCellID: 40, NeedNotItem: true}
	conn168.RecItemIDs = append(conn168.RecItemIDs, sword.ID, cloak.ID)
	// EAST
	conn169 := tbastory.MapConn{ID: 169, Name: "EAST", StartCellID: 49, EndCellID: 51}
	// WEST (if !sword)
	conn170 := tbastory.MapConn{ID: 170, Name: "WEST", StartCellID: 49, EndCellID: 45, NeedNotItem: true}
	conn170.RecItemIDs = append(conn170.RecItemIDs, sword.ID)
	// WEST (if sword)
	conn171 := tbastory.MapConn{ID: 171, Name: "WEST", StartCellID: 49, EndCellID: 48, NeedItem: true}
	conn171.RecItemIDs = append(conn171.RecItemIDs, sword.ID)
	// JOIN
	conn172 := tbastory.MapConn{ID: 172, Name: "JOIN", StartCellID: 49, EndCellID: 50}
	// Add conns
	docks.Conns = append(docks.Conns, conn166, conn167, conn168, conn169, conn170, conn171, conn172)

	// Ocean (51)
	// NORTH
	conn173 := tbastory.MapConn{ID: 173, Name: "NORTH", StartCellID: 51, EndCellID: 41}
	// EAST
	conn174 := tbastory.MapConn{ID: 174, Name: "EAST", StartCellID: 51, EndCellID: 52}
	// WEST
	conn175 := tbastory.MapConn{ID: 175, Name: "WEST", StartCellID: 51, EndCellID: 49}
	// Add conns
	village.Conns = append(village.Conns, conn173, conn174, conn175)

	// Cliff (52)
	// WEST
	conn176 := tbastory.MapConn{ID: 176, Name: "WEST", StartCellID: 52, EndCellID: 51}
	// Add conns
	village.Conns = append(village.Conns, conn176)

	// *****************
	// ** Cells (End) **
	// *****************

	// Orc Camp
	gygaria.Cells = append(gygaria.Cells, orcCamp)

	// Orc Camp (Empty)
	gygaria.Cells = append(gygaria.Cells, orcCampEmpty)

	// Castle Courtyard
	gygaria.Cells = append(gygaria.Cells, castleCourtyard)

	// Castle Courtyard (Seal)
	gygaria.Cells = append(gygaria.Cells, castleCourtyardSeal)

	// Castle
	gygaria.Cells = append(gygaria.Cells, castle)

	// Castle (Wizard Ending)
	gygaria.Cells = append(gygaria.Cells, castleEnd)

	// Land Slide
	gygaria.Cells = append(gygaria.Cells, landSlide)

	// Dense Forest
	gygaria.Cells = append(gygaria.Cells, forest)

	// Backwater Village
	gygaria.Cells = append(gygaria.Cells, village)

	// Backwater Village (Empty)
	gygaria.Cells = append(gygaria.Cells, villageEmpty)

	// Wolves' Den
	gygaria.Cells = append(gygaria.Cells, den)

	// Wolves' Den (Empty)
	gygaria.Cells = append(gygaria.Cells, denEmpty)

	// Spring Meadow
	gygaria.Cells = append(gygaria.Cells, meadow)

	// Spring Meadow (Seal)
	gygaria.Cells = append(gygaria.Cells, meadowSeal)

	// Spring Meadow (Empty)
	gygaria.Cells = append(gygaria.Cells, meadowEmpty)

	// Crossroads
	gygaria.Cells = append(gygaria.Cells, crossroads)

	// Crossroads (Deal)
	gygaria.Cells = append(gygaria.Cells, crossroadsDeal)

	// Crossroads (Empty)
	gygaria.Cells = append(gygaria.Cells, crossroadsEmpty)

	// Bridge
	gygaria.Cells = append(gygaria.Cells, bridge)

	// Boneyard
	gygaria.Cells = append(gygaria.Cells, boneyard)

	// Boneyard (Grab)
	gygaria.Cells = append(gygaria.Cells, boneyardGrab)

	// Boneyard (Empty)
	gygaria.Cells = append(gygaria.Cells, boneyardEmpty)

	// Boneyard (Thieves' Ending)
	gygaria.Cells = append(gygaria.Cells, boneyardEnd)

	// Fairy Glade
	gygaria.Cells = append(gygaria.Cells, glade)

	// Fairy Glade (Quest)
	gygaria.Cells = append(gygaria.Cells, gladeQuest)

	// Fairy Glade (Potion)
	gygaria.Cells = append(gygaria.Cells, gladePotion)

	// Crystal Lake
	gygaria.Cells = append(gygaria.Cells, lake)

	// Crystal Lake (Quest)
	gygaria.Cells = append(gygaria.Cells, lakeQuest)

	// Caravan
	gygaria.Cells = append(gygaria.Cells, caravan)

	// Plains
	gygaria.Cells = append(gygaria.Cells, plains)

	// Desert
	gygaria.Cells = append(gygaria.Cells, desert)

	// Desert (Empty)
	gygaria.Cells = append(gygaria.Cells, desertEmpty)

	// Grave Keeper's Hovel
	gygaria.Cells = append(gygaria.Cells, hovel)

	// Grave Keeper's Hovel (Deed Ending)
	gygaria.Cells = append(gygaria.Cells, hovelEnd)

	// Town Graveyard
	gygaria.Cells = append(gygaria.Cells, graveyard)

	// Town Graveyard (Inspect)
	gygaria.Cells = append(gygaria.Cells, graveyardInspect)

	// Town Graveyard (Empty)
	gygaria.Cells = append(gygaria.Cells, graveyardEmpty)

	// Town
	gygaria.Cells = append(gygaria.Cells, town)

	// Town (Old Man)
	gygaria.Cells = append(gygaria.Cells, townMan)

	// Town (Lady in Rags)
	gygaria.Cells = append(gygaria.Cells, townLady)

	// Farm
	gygaria.Cells = append(gygaria.Cells, farm)

	// Spider Forest
	gygaria.Cells = append(gygaria.Cells, spiderForest)

	// Spider Forest (Empty)
	gygaria.Cells = append(gygaria.Cells, spiderForestEmpty)

	// Mountain
	gygaria.Cells = append(gygaria.Cells, mountain)

	// Dragon's Lair
	gygaria.Cells = append(gygaria.Cells, dragonLair)

	// Dragon's Lair (Riddle)
	gygaria.Cells = append(gygaria.Cells, dragonLairRiddle)

	// Dragon's Lair (Sneak)
	gygaria.Cells = append(gygaria.Cells, dragonLairSneak)

	// Dragon's Lair (Empty)
	gygaria.Cells = append(gygaria.Cells, dragonLairEmpty)

	// Docks
	gygaria.Cells = append(gygaria.Cells, docks)

	// Docks (Sailor's Ending)
	gygaria.Cells = append(gygaria.Cells, docksEnd)

	// Ocean
	gygaria.Cells = append(gygaria.Cells, ocean)

	// Cliff
	gygaria.Cells = append(gygaria.Cells, cliff)

	return gygaria
}

// CreatePlayer returns a player set to default for Gygaria
func CreatePlayer() tbastory.Player {
	var player tbastory.Player

	player.Name = "Player"
	player.HP = 100
	player.Attack = 20
	player.Defense = 20
	player.CLocID = 39
	player.PLocID = 0

	return player
}
