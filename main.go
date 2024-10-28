package main

import (
	"fmt"
	"creatures"
)

func choose_creature() string {
	fmt.Println("Which creature do you choose ?")
	var creature string
	fmt.Scanln(&creature)
	return creature
}

func tab_creatures() []string {
	var list_creatures []string
	var dragon []string
	var griffin []string
	list_creatures = append(list_creatures, "Dragon", "Griffin", "Chimera", "Cerberus", "Cyclops", "Centaur", "Wendigo", "Gorgons", "Hydra", "Vampire", "Werewolf", "Basilica", "Sphinx", "Kraken", "Phoenix", "Leviathan", "Unicorn")
	fmt.Println(list_creatures)
	return list_creatures
}

func effects_special_attack(input string) string {
	var special_attack []string
	special_attack = append(special_attack, "poisoning", "confusion", "burn", "freeze", "encouragement", "blast", "flight")
	if input == "poisoning" {
		//enlèves 10 pv/tour pendant 5 tours
	}
	if input == "confusion" {
		//une chance sur 3 de ne plus être confus au prochain tour sinon reçois 10 dégats
	}
	if input == "burn" {
		//enlèves 25 pv/tour pendant 2 tours
	}
	if input == "freeze" {
		//une chance sur 4 que le joueur adverse soit bloqué pendant 2 tours
	}
	if input == "encouragement" {
		//une chance sur 3, attaques +50
	}
	if input == "blast" {
		//pv/3, 500 dégats à l'adversaire
	}
	if input == "flight" {
		//une chance sur deux, esquive le coup de l'adversaire au prochain tour
	}
	effect := ""
	return effect
}

func player_creature(player_num string) string {
	var input string
	var choice string
	run := true
	for run {
		fmt.Println(player_num + ", do you want to choose a creature (Reply choose) or look at a creature's abilities (Reply look)?")
		fmt.Scanln(&choice)
		if choice == "choose" {
			player_num = choose_creature()
			fmt.Println("\n")
			run = false
		} else if choice == "look" {
			fmt.Println("Which creature's abilities do you want to look")
			fmt.Scanln(&input)
			fmt.Println(creatures.input)
		} else {
			fmt.Println("Invalid choice")
		}
	}
	return player_num
}

func main() {
	tab_creatures() // Affiche la liste des créatures disponibles
	creature_j1 := player_creature("Player 1")
	creature_j2 := player_creature("Player 2")
	fmt.Println("Player 1, you chose " + creature_j1 + "\nPlayer 2, you chose " + creature_j2)
}
