package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"

	fmt.Println("Ї")
	// rune є типом даних, який представляє одиничний Unicode символ. Це тип даних, який є синонімом для int32
}
