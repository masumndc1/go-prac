package main

import (
	"fmt"
	"go-prac/shared"
)

func main() {
	var few_words = shared.Few_words
	var vowels = shared.Vowels
	fmt.Println(len(few_words))
	fmt.Println(len(vowels))
	//for i := 0; i < len(vowels); i++ {
	// modern way {
	for i := range len(vowels) {
		fmt.Printf("%d %s\n", i, vowels[i])
	}
	fmt.Println(vowels)
	fmt.Printf("3 element slices of vowels array %s\n", vowels[0:3])
}
