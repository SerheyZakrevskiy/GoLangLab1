package main

import "fmt"

func main() {
	var first, second bool
	var third bool = true
	fourth := !third
	var fifth = true

	fmt.Println("first  = ", first)       // false - не ініцілізували змінну first, go автоматично по дефолту оприділяє бульову змінну як false
	fmt.Println("second = ", second)      // false - не ініцілізували змінну second, go автоматично по дефолту оприділяє бульову змінну як false
	fmt.Println("third  = ", third)       // true   -  там явно вказане значення true
	fmt.Println("fourth = ", fourth)      // false !- цей оператор має логічний тип,third - true,але !third  перетворилось  на false, fourth - записала булеве значення -тому false
	fmt.Println("fifth  = ", fifth, "\n") // true - явно вказане значення

	fmt.Println("!true  = ", !true)        // false -  ! - оператор "не", тому false
	fmt.Println("!false = ", !false, "\n") // true  -  ! - оператор "не", тому true

	fmt.Println("true && true   = ", true && true)         // true - логічний оператор &&,якщо обидва true - то резульат буде true
	fmt.Println("true && false  = ", true && false)        // false  - логічний оператор &&,якщо якщо хоч один false - то резульат буде false
	fmt.Println("false && false = ", false && false, "\n") // false  -   логічний оператор &&,якщо якщо обидва false - то резульат буде false

	fmt.Println("true || true   = ", true || true)         // true - логічний оператор ||,якщо обидва true - то резульат буде true
	fmt.Println("true || false  = ", true || false)        // true  - логічний оператор ||,якщо якщо хоч один true - то резульат буде true
	fmt.Println("false || false = ", false || false, "\n") // false  - логічний оператор ||якщо якщо хоч один false - то резульат буде false

	fmt.Println("2 < 3  = ", 2 < 3)        // true  - логічний оператор <,в нашому випадку true,тому що 2 меньше за 3
	fmt.Println("2 > 3  = ", 2 > 3)        // false  - логічний оператор <,в нашому випадку false,тому що 2 меньше за 3
	fmt.Println("3 < 3  = ", 3 < 3)        // false  - логічний оператор <,в нашому випадку false,тому що 3 рівно 3
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true   -  логічний оператор <=,в нашому випадку true,тому що 3 меньше-рівно 3
	fmt.Println("3 > 3  = ", 3 > 3)        // false  -логічний оператор <,в нашому випадку false,тому що 3 рівно 3
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true    логічний оператор >=,в нашому випадку true,тому що 3 більше-рівно 3
	fmt.Println("2 == 3 = ", 2 == 3)       // false логічний оператор ==,в нашому випадку false,тому що 2 не рівно 3
	fmt.Println("2 != 3 = ", 2 != 3)       // true - логічний оператор !=,в нашому випадку true,тому що 2 не рівно 3
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false  -  логічний оператор !=,в нашому випадку false,тому що 3 рівно 3

	//Задание.
	//1. Пояснить результаты операций
}
