package main

import "fmt"

func main() {
	// countries := make(map[string]string)
	countries := make(map[string]string, 5)

	countries["Mexico"] = "D.F"
	countries["Cuba"] = "Havana"
	countries["Argentina"] = "Buenos Aires"

	// fmt.Println(countries)
	// fmt.Println(countries)
	// fmt.Println(countries["Cuba"])

	champion := map[string]int{
		"Barcelona":    20,
		"Real Madrid":  40,
		"Chivas":       30,
		"Boca Juniors": 67,
	}

	fmt.Println(champion)

	champion["Atletic"] = 67
	champion["Chivas"] = 98
	fmt.Println(champion)

	delete(champion, "Atletic")
	fmt.Println(champion)

	for team, value := range champion {
		fmt.Printf("The team %s, has a value of %d\n", team, value)
	}

	value, exists := champion["Mineiro"]
	fmt.Printf("El puntaje capturado  es %d y el equipo existe %t\n", value, exists)

	value1, exists1 := champion["Chivas"]
	fmt.Printf("El puntaje capturado  es %d y el equipo existe %t\n", value1, exists1)

}
