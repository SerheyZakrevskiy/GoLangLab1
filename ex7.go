package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	//Задание.
	//1. Создайте 2 переменные  разных типов. Выпоните арифметические операции. Результат вывести

	num1 := 23
	num2 := 2.5
	result1 := float64(num1) / num2
	result2 := float64(num1) * num2
	result3 := float64(num1) + num2
	fmt.Printf("Перший приклад: %d / %f = %.2f\n", num1, num2, result1)
	fmt.Printf("Другий приклад: %d / %f = %.2f\n", num1, num2, result2)
	fmt.Printf("Третій приклад: %d / %f = %.2f\n", num1, num2, result3)

}
