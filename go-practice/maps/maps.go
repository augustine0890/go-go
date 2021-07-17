package main

import "fmt"

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	var mp map[string]int = map[string]int{
		"apple":  5,
		"pear":   7,
		"orange": 8,
	}

	mp["apple"] = 9
	mp["augustine"] = 434
	fmt.Println(mp["apple"])
	fmt.Println(mp)
	fmt.Println(mp["apple"])

	delete(mp, "augustine")
	fmt.Println(len(mp))

	val, ok := mp["augustine"]
	fmt.Println(val, ok)

	// mp2 := make(map[string]int)
	// fmt.Println(mp2)

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Taxas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Florida"])
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	_, ok = statePopulations["Flo"]
	fmt.Println(ok)

	aDoctor := Doctor{
		Number:    3,
		ActorName: "John Pertwee",
		Companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.Companions)
}
