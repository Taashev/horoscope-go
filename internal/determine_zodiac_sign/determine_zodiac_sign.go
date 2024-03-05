package determine_zodiac_sign

// Определяет знак зодиака
func DetermineZodiacSign(day, month int) string {
	switch {
	case (month == 3 && day >= 21) || (month == 4 && day <= 19):
		return "Овен"
	case (month == 4 && day >= 20) || (month == 5 && day <= 20):
		return "Телец"
	case (month == 5 && day >= 21) || (month == 6 && day <= 20):
		return "Близнецы"
	case (month == 6 && day >= 21) || (month == 7 && day <= 22):
		return "Рак"
	case (month == 7 && day >= 23) || (month == 8 && day <= 22):
		return "Лев"
	case (month == 8 && day >= 23) || (month == 9 && day <= 22):
		return "Дева"
	case (month == 9 && day >= 23) || (month == 10 && day <= 22):
		return "Весы"
	case (month == 10 && day >= 23) || (month == 11 && day <= 21):
		return "Скорпион"
	case (month == 11 && day >= 22) || (month == 12 && day <= 21):
		return "Стрелец"
	case (month == 12 && day >= 22) || (month == 1 && day <= 19):
		return "Козерог"
	case (month == 1 && day >= 20) || (month == 2 && day <= 18):
		return "Водолей"
	case (month == 2 && day >= 19) || (month == 3 && day <= 20):
		return "Рыбы"
	default:
		return "Не удалось определить знак зодиака"
	}
}
