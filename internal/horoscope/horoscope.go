package horoscope

import "fmt"

func ZodiacEnglishMapping(zodiac string) (string, error) {
	englishZodiac, ok := map[string]string{
		"Овен":     "aries",
		"Телец":    "taurus",
		"Близнецы": "gemini",
		"Рак":      "cancer",
		"Лев":      "leo",
		"Дева":     "virgo",
		"Весы":     "libra",
		"Скорпион": "scorpio",
		"Стрелец":  "sagittarius",
		"Козерог":  "capricorn",
		"Водолей":  "aquarius",
		"Рыбы":     "pisces",
	}[zodiac]

	if !ok {
		return "", fmt.Errorf("неизвестный знак зодиака: %s", zodiac)
	}

	return englishZodiac, nil
}
