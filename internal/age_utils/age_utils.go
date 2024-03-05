package ageutils

import (
	"fmt"
	"time"
)

func GetYearOfBirth() int {
	var yearOfBirth int

	fmt.Scanln(&yearOfBirth)

	return yearOfBirth
}

func GetMonthOfBirth() int {
	var monthOfBirth int

	fmt.Scanln(&monthOfBirth)

	return monthOfBirth
}

func GetDayOfBirth() int {
	var dayOfBirth int

	fmt.Scanln(&dayOfBirth)

	return dayOfBirth
}

func PluralizeYears(age int) string {
	if age%10 == 1 && age != 11 {
		return "год"
	} else if (age%10 == 2 || age%10 == 3 || age%10 == 4) && (age < 10 || age > 20) {
		return "года"
	}
	return "лет"
}

func CalculateAge(yearOfBirth int) int {
	// Получим текущий год
	currentYear := time.Now().Year()

	// Рассчитаем возраст пользователя
	return currentYear - yearOfBirth
}
