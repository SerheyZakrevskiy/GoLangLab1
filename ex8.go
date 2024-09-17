package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	//Задание.
	//1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести
	num1 := 23
	num2 := 2.5
	result1 := float64(math.MaxFloat64) / num2
	result2 := float64(num1) * math.Pi
	fmt.Printf("Перший приклад: %d / %f = %.2f\n", num1, math.MaxFloat64 , result1)
	fmt.Printf("Другий приклад: %f / %f = %.2f\n", math.Pi , num2, result2)
	fmt.Printf("Третій приклад: %d / %f = %.2f\n", num1, num2, result3)



}
