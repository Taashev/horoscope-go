package greeting

import "fmt"

func PrintGreeting(name string, age int, yearsSuffix, zodiacSign string) {
	fmt.Printf("Привет, %s! Вам %d %s. Ваш знак зодиака - %s.\n", name, age, yearsSuffix, zodiacSign)
}
