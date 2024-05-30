package main

import "fmt"

func main() {

	fmt.Println("Task 8.1")
	m1 := map[string]struct{}{ //& 8.3
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	fmt.Println(m1)

	fmt.Println("Task 8.2")
	s := []string{"слон", "бегемот", "осьминог"}
	m2 := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1} // &8.5
	for _, sv := range s {
		mk, ok := m2[sv]
		if ok {
			fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", sv, mk, m2)
		} else { // not sure it's neсessary
			fmt.Printf("Животное: %s (отсутствует в карте: %v)\n", sv, m2)
		}
	}

	fmt.Println("Task 8.3")
	//m1 from 8.1
	delete(m1, "бегемот")
	fmt.Println(m1)

	fmt.Println("Task 8.4")
	m4 := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	m4["выдра"] = struct{}{}
	fmt.Println(m4)

	fmt.Println("Task 8.5")
	//m2 from 8.2
	m2["бегемот"] = 2
	fmt.Println(m2)
}
