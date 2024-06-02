package main

import "fmt"

func fruitMarket(s string) int { //9.1
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

func checkFood(s string) { // 9.3
	switch s {
	case "апельсин", "груша", "яблоко":
		fmt.Println("фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("овощ")
	default:
		fmt.Println("что-то странное...")
	}

}

func main() {
	fmt.Println("Task 9.1")
	s1 := []string{"апельсины", "мандарины", "яблоки"} //slice of fruits
	for _, v := range s1 {
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

	fmt.Println("Task 9.3")
	s3 := []string{"апельсин", "мандарин", "помидор", "томат"} //slice of food
	for _, v := range s3 {
		fmt.Print(v, ": ")
		checkFood(v)
	}
}
