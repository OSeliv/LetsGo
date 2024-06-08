package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() (string, error) //3rd way + error
}
type Duck struct {
	voice string
}

// ----- 3rd way --------
func (d *Duck) Sing() (string, error) {
	// adding this block
	if d == nil { // checking for nil
		return "That's all folks", errors.New("Broken voice!")
	}
	// --------
	return d.voice, nil
}

// -------------------------

func main() {
	var song string
	// ------ 1st way -----
	fmt.Println("1st solution")
	// This is error == var d *Duck ==  nil pointer to Duck (not an object Duck) without value, Sing() not expects that
	//	var d *Duck
	//	song, err := Sing(d)
	var d1 = Duck{"quack-quack"} // it can be just "var d1 Duck"
	song, err := Sing(&d1)       // pointer to d1

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
	printSong(song)

	// ----- 2nd way --------
	fmt.Println()
	fmt.Println("2nd solution")
	//	var d *Duck
	//	song, err := Sing(d)
	d2 := &Duck{"quack"} // creating object Duck
	song, err = Sing(d2) // passing through the pointer d2

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
	printSong(song)

	// ----- 3rd way --------
	fmt.Println()
	fmt.Println("3rd solution")
	var d3 *Duck         // unchanged
	song, err = Sing(d3) // unchanged

	if err != nil {
		fmt.Println(err)
		fmt.Println(song)
		return
	}
	printSong(song)
}

func Sing(b Bird) (string, error) {
	/* // was changed for 3rd way
	if b != nil {
		return b.Sing(), nil
	}*/
	return b.Sing() //, errors.New("Ошибка пения!")  // was changed for 3rd way
}

func printSong(s string) {
	fmt.Println("\nOld MacDonald had a farm, E-I-E-I-O.\nAnd on that farm he had a duck, E-I-E-I-O.")
	fmt.Printf("With a '%s' here and a '%s' there.\nOld MacDonald had a farm, E-I-E-I-O.\n", s, s)
}
