package main

import "fmt"

func fruitMarket(s string) int {
	fruits := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	v, ok := fruits[s]
	if ok {
		return v
	} else {
		fmt.Println("отсутствует в ассортименте")
		return -1
	}
}

func main() {
	fmt.Println("Task 9.1")
	s := []string{"апельсины", "мандарины", "яблоки"} //slice of fruits
	for _, v := range s {
		fmt.Printf("%s - ", v)
		val := fruitMarket(v)
		if val >= 0 {
			fmt.Println(val)
		}
	}

	fmt.Println("Task 9.2")
	sl := []int{1, 2, 3}

	for _, c1 := range sl { //cycle1
		fmt.Println("v1:", c1)
	OUT:
		for _, c2 := range sl { //cycle2
			fmt.Println("\tv2:", c2)
			for _, c3 := range sl { //cycle3
				fmt.Println("\t\tv3:", c3)
				for i, c4 := range sl { //cycle4
					fmt.Println("\t\t\tv4:", c4)
					if i == 1 {
						break OUT
					}
				}
			}
		}
	}
}
