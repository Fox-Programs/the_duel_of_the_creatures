package piscine

// Données de la créature
	// points de vie
	// capcité spéciale de la créature
	// attaque de base 1
	// attaque de base 2
	// attaque spéciale 1
	// attaque spéciale 2
	// latence pour pouvoir utiliser l'attaque 1
	// latence pour pouvoir utiliser l'attaque spéciale 1
	// latence pour pouvoir utiliser l'attaque spéciale2
	
type DragonData struct {
	pv := 1000
	capacity := "damage received minimized"
	basic_attack_1 := 40
	basic_attack_2 := 70
	special_attack_1 := "poisoning"
	special_attack_2 := "confusion"
	latency_attack_2 := 3
	latency_special_attack_1 := 3
	latency_special_attack_2 := 3
}

type GriffinData struct {
	pv := 800
	capacity := "Chance that the attack received has no effect"
	basic_attack_1 := 60
	basic_attack_2 := 90
	special_attack_1 := "poisoning"
	special_attack_2 := "flight"
	latency_attack_2 := 3
	latency_special_attack_1 := 3
	latency_special_attack_2 := 3
}
	