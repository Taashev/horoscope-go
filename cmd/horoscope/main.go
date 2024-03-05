package run

import (
	"fmt"
	horoscope "horoscope/api/horoscope"
	ageUtils "horoscope/internal/age_utils"
	getZodiac "horoscope/internal/determine_zodiac_sign"
	getName "horoscope/internal/get_name"
	greeting "horoscope/internal/greeting"
)

func Run() {
	fmt.Print("Введите ваше имя: ")
	name := getName.GetName()

	fmt.Print("Введите год вашего рождения: ")
	yearOfBirth := ageUtils.GetYearOfBirth()

	fmt.Print("Введите месяц вашего рождения (число): ")
	monthOfBirth := ageUtils.GetMonthOfBirth()

	fmt.Print("Введите день вашего рождения: ")
	dayOfBirth := ageUtils.GetDayOfBirth()

	// Рассчитаем возраст пользователя
	age := ageUtils.CalculateAge(yearOfBirth)

	// Получим правильное склонение для слова "год"
	yearsSuffix := ageUtils.PluralizeYears(age)

	// Определим знак зодиака
	zodiacSign := getZodiac.DetermineZodiacSign(dayOfBirth, monthOfBirth)

	// Получим гороскоп
	horoscope, err := horoscope.GetHoroscope(zodiacSign)

	if err != nil {
		fmt.Printf("Не удалось получить гороскоп: %v\n", err)
		return
	}

	greeting.PrintGreeting(name, age, yearsSuffix, zodiacSign)

	fmt.Println("Гороскоп на сегодня:")
	fmt.Println(horoscope)
}
